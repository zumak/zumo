package main

import (
	"context"
	"crypto/sha1"
	"io"
	"net"
	"os"
	"path/filepath"

	"github.com/codingconcepts/env"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Config struct {
	Bind   string `env:"bind" required:"true"`
	DBPath string `env:"db-path" required:"true"`
}

type FileInfo struct {
	gorm.Model
	SHA1  string
	Path  string
	Error string
	Size  uint64
}

type Server struct {
	db *gorm.DB
}

func main() {
	config := &Config{}
	if err := env.Set(config); err != nil {
		logrus.Fatal(err)
	}

	logrus.Debugf("config: %+v", config)

	db, err := gorm.Open("sqlite3", config.DBPath)
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&FileInfo{}).Error; err != nil {
		panic(err)
	}

	server := &Server{
		db,
	}
	go server.deamon()

	// TODO grpc
	lis, err := net.Listen("tcp", config.Bind)
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer(
	//grpc.StreamInterceptor(utils.StreamMetadataBypass),
	//grpc.UnaryInterceptor(utils.UnaryMetadataBypass),
	)
	// ********** Register grpc Server ************ //
	// datatype.RegisterClusterManagerServer(grpcServer, &ClusterManagerServer{server})
	// determine whether to use TLS

	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
}
func (server *Server) deamon() {

}

func scan(ctx context.Context, path string) error {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		return nil

		// JUST scan..
		// TODO put queue?
		fi := &FileInfo{
			Path: path,
		}

		server
		return nil

	})
	if err != nil {
		return err
	}
	return nil
}
func calcSHA1(ctx context.Context, path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := sha1.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	return string(h.Sum(nil)), nil
}
func (server *Server) startCalcSHA1(ctx context.Context) {
	for {
		info := &FileInfo{}
		// sha1 not exist & error not exist
		err := server.db.Model(&FileInfo{}).First(info).Error
		hash, err := calcSHA1(ctx, info.Path)
		if err != nil {
			// write error
			info.Error = err.Error()
		} else {
			info.SHA1 = hash
		}
		server.db.Save(info)
	}
}
