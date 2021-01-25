package main

import (
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"time"
)

var (
	passwd string = "DFwl123456"
	user   string = "root"
	host   string = "18.20.6.8"
	port   int    = 22
)

//func connect0(user, password, host string, port int) (*ssh.Session, error) {
//
//	var (
//		addr         string
//		clientConfig *ssh.ClientConfig
//		client       *ssh.Client
//		session      *ssh.Session
//		err          error
//	)
//
//	clientConfig = &ssh.ClientConfig{
//		User: user,
//		Auth: []ssh.AuthMethod{
//			ssh.Password(password),
//		},
//		Timeout:         30 * time.Second,
//		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
//		//HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
//		//	return nil
//		//},
//	}
//
//	// connet to ssh
//	addr = fmt.Sprintf("%s:%d", host, port)
//	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
//		return nil, err
//	}
//
//	// create session
//	if session, err = client.NewSession(); err != nil {
//		return nil, err
//	}
//
//	return session, nil
//
//}

var localPath = "E:/work/项目支持/上海/松江项目/松江升级/imageRequestProxyNonMysql"
var remotePath = "/home/"

func main() {
	servers := initArray()
	CopyFile(servers)
}

type Server struct {
	Host     string
	Port     int
	UserName string
	Password string
}

func initArray() []*Server {
	servers := make([]*Server, 0)
	//servers = append(servers, &Server{Host: "18.20.6.5", Port: 22, UserName: "root", Password: "DFwl123456",})
	//servers = append(servers, &Server{Host: "18.20.6.7", Port: 22, UserName: "root", Password: "DFwl123456",})
	//servers = append(servers, &Server{Host: "18.20.6.4", Port: 22, UserName: "root", Password: "DFwl123456",})
	//servers = append(servers, &Server{Host: "18.20.6.6", Port: 22, UserName: "root", Password: "DFwl123456",})
	//servers = append(servers, &Server{Host: "18.20.18.2 ", Port: 22, UserName: "root", Password: "DFwl123456",})

	//TODO 复制失败,已手动复制
	// servers = append(servers, &Server{Host: "18.20.7.2", Port: 22, UserName: "root", Password: "DFwl123456",})
	//servers = append(servers, &Server{Host: "18.20.10.6", Port: 22, UserName: "root", Password: "DFwl123456",})
	//TODO 复制失败，已手动复制
	// servers = append(servers, &Server{Host: "18.20.7.10", Port: 22, UserName: "root", Password: "DFwl123456",})

	//servers = append(servers, &Server{Host: "18.20.9.11", Port: 22, UserName: "root", Password: "DFwl123456",})
	//servers = append(servers, &Server{Host: "18.3.0.101", Port: 22, UserName: "root", Password: "DFwl123456",})
	//servers = append(servers, &Server{Host: "18.17.10.216", Port: 22, UserName: "root", Password: "DFwl123456",})
	//servers = append(servers, &Server{Host: "18.17.10.236", Port: 22, UserName: "root", Password: "DFwl123456",})
	//servers = append(servers, &Server{Host: "18.17.11.20", Port: 22, UserName: "root", Password: "DFwl123456",})
	//servers = append(servers, &Server{Host: "18.17.17.83", Port: 22, UserName: "root", Password: "DFwl123456",})
	//servers = append(servers, &Server{Host: "18.20.4.7", Port: 22, UserName: "root", Password: "DFwl123456",})
	//servers = append(servers, &Server{Host: "18.8.81.100", Port: 22, UserName: "root", Password: "DFwl123456",})

	//TODO 复制失败，已手动复制
	// servers = append(servers, &Server{Host: "18.20.11.2", Port: 22, UserName: "root", Password: "DFwl123456",})
	//TODO 复制失败，已手动复制
	// servers = append(servers, &Server{Host: "18.20.7.14", Port: 22, UserName: "root", Password: "DFwl123456",})

	//servers = append(servers, &Server{Host: "18.20.11.1", Port: 22, UserName: "root", Password: "DFwl123456",})
	//servers = append(servers, &Server{Host: "18.20.3.21", Port: 22, UserName: "root", Password: "DFwl123456",})
	//servers = append(servers, &Server{Host: "18.20.3.23", Port: 22, UserName: "root", Password: "DFwl123456",})
	//servers = append(servers, &Server{Host: "18.20.3.22", Port: 22, UserName: "root", Password: "DFwl123456",})
	//servers = append(servers, &Server{Host: "18.20.15.1", Port: 22, UserName: "root", Password: "DFwl123456",})
	//servers = append(servers, &Server{Host: "18.20.6.2", Port: 22, UserName: "root", Password: "DFwl123456",})
	//servers = append(servers, &Server{Host: "18.20.16.1", Port: 22, UserName: "root", Password: "DFwl123456",})
	//servers = append(servers, &Server{Host: "18.20.4.5", Port: 22, UserName: "root", Password: "DFwl123456",})
	//servers = append(servers, &Server{Host: "18.20.10.5", Port: 22, UserName: "root", Password: "DFwl123456",})
	//servers = append(servers, &Server{Host: "18.20.10.4", Port: 22, UserName: "root", Password: "DFwl123456",})
	//servers = append(servers, &Server{Host: "18.20.9.2", Port: 22, UserName: "root", Password: "DFwl123456",})
	//servers = append(servers, &Server{Host: "18.20.14.1", Port: 22, UserName: "root", Password: "DFwl123456",})
	//servers = append(servers, &Server{Host: "18.20.10.2", Port: 22, UserName: "root", Password: "DFwl123456",})
	//servers = append(servers, &Server{Host: "18.20.4.2", Port: 22, UserName: "root", Password: "DFwl123456",})
	//servers = append(servers, &Server{Host: "18.11.8.174", Port: 22, UserName: "root", Password: "DFwl123456",})
	//servers = append(servers, &Server{Host: "18.20.4.4", Port: 22, UserName: "root", Password: "123456",})
	//servers = append(servers, &Server{Host: "18.20.4.3", Port: 22, UserName: "root", Password: "DFwl123456",})
	//servers = append(servers, &Server{Host: "18.20.6.1", Port: 22, UserName: "root", Password: "DFwl123456",})
	//servers = append(servers, &Server{Host: "18.20.1.1", Port: 22, UserName: "root", Password: "DFwl123456",})
	//servers = append(servers, &Server{Host: "18.20.5.1", Port: 22, UserName: "root", Password: "DFwl123456",})
	//servers = append(servers, &Server{Host: "18.3.198.50", Port: 22, UserName: "root", Password: "DFwl123456",})

	//TODO 网络不通
	// servers = append(servers, &Server{Host: "18.20.18.1", Port: 22, UserName: "root", Password: "123456",})
	// servers = append(servers, &Server{Host: "18.20.0.1", Port: 22, UserName: "root", Password: "DFwl123456",})

	return servers
}

func CopyFile(servers []*Server) {
	for _, server := range servers {
		DoBackup(strings.TrimSpace(server.Host), server.Port, server.UserName, server.Password, localPath, remotePath)
		time.Sleep(1 * time.Second)
	}

}

func DoBackup(host string, port int, userName string, password string, localPath string, remotePath string) {
	var (
		err        error
		sftpClient *sftp.Client
	)
	start := time.Now()
	sftpClient, err = connect(userName, password, host, port)
	if err != nil {
		log.Println(host, err)
	}
	defer sftpClient.Close()

	_, errStat := sftpClient.Stat(remotePath)
	if errStat != nil {
		log.Println(remotePath + " remote path not exists!")
	}

	_, err = ioutil.ReadDir(localPath)
	if err != nil {
		log.Println(localPath + " local path not exists!")
	}
	uploadDirectory(sftpClient, localPath, remotePath)
	elapsed := time.Since(start)
	fmt.Println("elapsed time : ", elapsed)
}

func connect(user, password, host string, port int) (*sftp.Client, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		sshClient    *ssh.Client
		sftpClient   *sftp.Client
		err          error
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))

	clientConfig = &ssh.ClientConfig{
		User:            user,
		Auth:            auth,
		Timeout:         30 * time.Second,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), //ssh.FixedHostKey(hostKey),
	}

	// connet to ssh
	addr = fmt.Sprintf("%s:%d", host, port)
	if sshClient, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}

	// create sftp client
	if sftpClient, err = sftp.NewClient(sshClient); err != nil {
		return nil, err
	}
	return sftpClient, nil
}

//上传文件夹
func uploadDirectory(sftpClient *sftp.Client, localPath string, remotePath string) {
	localFiles, err := ioutil.ReadDir(localPath)
	if err != nil {
		log.Println("read dir list fail ", err)
	}

	for _, backupDir := range localFiles {
		localFilePath := path.Join(localPath, backupDir.Name())
		remoteFilePath := path.Join(remotePath, backupDir.Name())
		if backupDir.IsDir() {
			sftpClient.Mkdir(remoteFilePath)
			uploadDirectory(sftpClient, localFilePath, remoteFilePath)
		} else {
			uploadFile(sftpClient, path.Join(localPath, backupDir.Name()), remotePath)
		}
	}

	fmt.Println(localPath + "  copy directory to remote server finished!")
}

//上传文件
func uploadFile(sftpClient *sftp.Client, localFilePath string, remotePath string) {
	srcFile, err := os.Open(localFilePath)
	if err != nil {
		fmt.Println("os.Open error : ", localFilePath)
		log.Fatal(err)

	}
	defer srcFile.Close()

	var remoteFileName = path.Base(localFilePath)

	dstFile, err := sftpClient.Create(path.Join(remotePath, remoteFileName))
	if err != nil {
		fmt.Println("sftpClient.Create error : ", path.Join(remotePath, remoteFileName))
		log.Fatal(err)

	}
	defer dstFile.Close()

	ff, err := ioutil.ReadAll(srcFile)
	if err != nil {
		fmt.Println("ReadAll error : ", localFilePath)
		log.Fatal(err)

	}
	dstFile.Write(ff)
	fmt.Println(localFilePath + "  copy file to remote server finished!")
}
