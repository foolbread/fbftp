/*
author: foolbread
file: protocol/protocol.go
date: 2017/9/11
*/
package protocol

import (
	"github.com/foolbread/fbcommon/golog"
	"github.com/foolbread/fbftp/session"
	"github.com/foolbread/fbftp/con"
	"errors"
	"strings"
)

func InitProtocol(){
	golog.Info("fbftp protocol initing......")
}

const(
	FTP_DATACONN        =  150
	FTP_NOOPOK          = 200
	FTP_TYPEOK          = 200
	FTP_PORTOK          = 200
	FTP_EPRTOK          = 200
	FTP_UMASKOK         = 200
	FTP_CHMODOK         = 200
	FTP_EPSVALLOK       = 200
	FTP_STRUOK          = 200
	FTP_MODEOK          = 200
	FTP_PBSZOK          = 200
	FTP_PROTOK          = 200
	FTP_OPTSOK          = 200
	FTP_ALLOOK          = 202
	FTP_FEAT            = 211
	FTP_STATOK          = 211
	FTP_SIZEOK          = 213
	FTP_MDTMOK          = 213
	FTP_STATFILE_OK     = 213
	FTP_SITEHELP        = 214
	FTP_HELP            = 214
	FTP_SYSTOK          = 215
	FTP_GREET           = 220
	FTP_GOODBYE         = 221
	FTP_ABOR_NOCONN     = 225
	FTP_TRANSFEROK      = 226
	FTP_ABOROK          = 226
	FTP_PASVOK          = 227
	FTP_EPSVOK          = 229
	FTP_LOGINOK         = 230
	FTP_AUTHOK          = 234
	FTP_CWDOK           = 250
	FTP_RMDIROK         = 250
	FTP_DELEOK          = 250
	FTP_RENAMEOK        = 250
	FTP_MD5OK           = 251
	FTP_PWDOK           = 257
	FTP_MKDIROK         = 257
	FTP_GIVEPWORD       = 331
	FTP_RESTOK          = 350
	FTP_RNFROK          = 350
	FTP_IDLE_TIMEOUT    = 421
	FTP_DATA_TIMEOUT    = 421
	FTP_TOO_MANY_USERS  = 421
	FTP_IP_LIMIT        = 421
	FTP_IP_DENY         = 421
	FTP_TLS_FAIL        = 421
	FTP_BADSENDCONN     = 425
	FTP_BADSENDNET      = 426
	FTP_BADSENDFILE     = 451
	FTP_BADCMD          = 500
	FTP_BADOPTS         = 501
	FTP_COMMANDNOTIMPL  = 502
	FTP_NEEDUSER        = 503
	FTP_NEEDRNFR        = 503
	FTP_BADPBSZ         = 503
	FTP_BADPROT         = 503
	FTP_BADSTRU         = 504
	FTP_BADMODE         = 504
	FTP_BADAUTH         = 504
	FTP_NOSUCHPROT      = 504
	FTP_NEEDENCRYPT     = 522
	FTP_EPSVBAD         = 522
	FTP_DATATLSBAD      = 522
	FTP_LOGINERR        = 530
	FTP_NOHANDLEPROT    = 536
	FTP_FILEFAIL        = 550
	FTP_NOPERM          = 550
	FTP_UPLOADFAIL      = 553
)

var commandMap map[string]Command = map[string]Command{
	"ABOR":&commandAbor{},
	"\377\364\377\362ABOR":&commandAbor{},
	"CDUP":&commandCdup{},
	"CWD" :&commandCwd{},
	"DELE":&commandDele{},
	"FEAT":&commandFeat{},
	"HELP":&commandHelp{},
	"LIST":&commandList{},
	"MKD" :&commandMkd{},
	"NOOP":&commandNoop{},
	"PASS":&commandPass{},
	"PASV":&commandPasv{},
	"PWD" :&commandPwd{},
	"QUIT":&commandQuit{},
	"RETR":&commandRetr{},
	"RNFR":&commandRnfr{},
	"RNTO":&commandRnto{},
	"RMD" :&commandRmd{},
	"SIZE":&commandSize{},
	"STAT":&commandStat{},
	"STOR":&commandStor{},
	"SYST":&commandSyst{},
	"TYPE":&commandType{},
	"USER":&commandUser{},
}

var(
	errUnknowCmd = errors.New("unknow cmd error")
	errArgMiss = errors.New("arg is nil")
	errAuth = errors.New("auth error")
)


type Command interface {
	IsExtend() bool
	RequireAuth() bool
	RequireParam() bool
	Execute(*session.FTPSession, string)error
}

func DisPatchCmd(sess *session.FTPSession,req *con.FTPCmdReq)error{
	e := commandMap[strings.ToUpper(req.Cmd)]
	if e == nil{
		sess.CtrlCon.WriteMsg(FTP_BADCMD,"Unknown command.")
		return errUnknowCmd
	}

	if e.RequireParam()&& req.IsArgNULL(){
		return errArgMiss
	}

	if e.RequireAuth() && !sess.IsLogin(){
		return errAuth
	}

	return e.Execute(sess,req.Arg)
}
