#
# クライアントからサーバにセーブするマン
#

extends Node

signal _on_server_response

#セーブする関数
func save_server(userid, key, val) -> bool:
	_save_server.rpc(JSON.stringify({"userid": userid,"k": key,"v": val}))
	var ret = await _on_server_response
	# ここにエラー処理を書いたほうがいい
	return false
	
#ロードする関数
func load_server(userid, key) -> Dictionary:
	_load_server.rpc(JSON.stringify({"userid": userid,"k": key}))
	var ret = await _on_server_response
	# ここにエラー処理を書いたほうがいい
	return ret

# サーバ側のセーブ処理
@rpc("any_peer")
func _save_server(json_data):
	var data = JSON.parse_string(json_data)
	$HTTPRequest.request("http://localhost:6501/save/", [], HTTPClient.METHOD_POST, JSON.stringify(data))

# サーバ側のロード処理
@rpc("any_peer")
func _load_server(json_data):
	var data = JSON.parse_string(json_data)
	$HTTPRequest.request("http://localhost:6501/save/", [], HTTPClient.METHOD_GET, JSON.stringify(data))

# HTTPリクエストの返答があったときシグナルにクライアントに返す処理を接続
func _ready():
	$HTTPRequest.request_completed.connect(func(result, response_code, headers, body): 
		# GodotのServer側 <- Golang <- データベース
		var text = body.get_string_from_utf8()
		
		# 送ってきたやつに返すRPC
		_on_load_response.rpc_id( multiplayer.get_remote_sender_id(), text)
	)
	
# クライアント側の反応
@rpc()
func _on_load_response(text):
	var data = JSON.parse_string(text)
	_on_server_response.emit(data)
	
	
