package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	bucketName := "teste-alunos-chico"
	filePath := "data.json"

	key := filepath.Base(filePath)

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("Erro ao carregar configuração AWS: %v", err)
	}

	client := s3.NewFromConfig(cfg)

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Erro ao abrir o arquivo %v", err)
	}

	defer file.Close()

	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
		Body:   file,
	})
	if err != nil {
		log.Fatalf("Erro ao enviar o arquivo ao S3 bucket: %v", err)
	}
	fmt.Println("Arquivo enviado com sucesso.")

}
