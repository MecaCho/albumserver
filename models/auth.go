package models


func Login(user,password string)(bool,error){
	if user == "qwq" && password == "123456"{
		return true,nil
	}
	return false,nil
}
