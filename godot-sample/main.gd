extends Node2D

signal _on_server_response

@onready var saveman = $SaveMan # これでセーブするよ

var PORT = 1113 # 部屋のドア番号
var SERVER_IP = "localhost" #部屋の住所
var SERVER_URL = "localhost" #部屋の住所

#限りなく適当な表示
func print2(text):
	$Control/Panel/Kekka.text = str(text)

# 部屋にあそびにいく (クライアント）おした
func _on_client_pressed() -> void:
	if multiplayer.get_peers().size() == 0:
		var peer = WebSocketMultiplayerPeer.new()
		
		SERVER_URL = SERVER_IP+":"+str(PORT)
		
		peer.create_client(SERVER_URL)
		multiplayer.multiplayer_peer = peer
		
		await multiplayer.connected_to_server #接続完了まで待つ
		
		print2("サーバにつながった")

# 部屋をつくる （サーバー）おした
func _on_room_pressed() -> void:
	var peer = WebSocketMultiplayerPeer.new()
	peer.create_server(PORT)
	
	multiplayer.multiplayer_peer = peer
	
	multiplayer.peer_connected.connect(on_new_player)
	multiplayer.peer_disconnected.connect(on_exit_player)
	
	print2("サーバになった")
	var nodes = $Control/Panel.get_children()
	for n in nodes:
		if n.name == "Kekka":
			continue
		n.visible = false
	
func on_new_player(player_id):
	print2("だれかあそびにきたよ プレイヤー番号:"+ str(player_id))
	
func on_exit_player(player_id):
	print2("かえったよ プレイヤー番号: "+ str(player_id) )

func _on_createuser_pressed() -> void:
	pass

# セーブボタンおした
func _on_save_pressed() -> void:
	var userid = $Control/Panel/Documentid.text
	var key = $Control/Panel/Key.text
	var value = $Control/Panel/Value.text
	
	var err = await saveman.save_server(userid, key, value) #セーブする documentid + key + value
	
# ロードボタンおした
func _on_load_pressed() -> void:
	var userid = $Control/Panel/Documentid.text
	var key = $Control/Panel/Key.text
	
	var res = await saveman.load_server(userid, key) #ロードする documentid + key 
	#サーバの返事がDictonaryでresに帰ってくる
	print2("返事:"+ JSON.stringify(res))

	
