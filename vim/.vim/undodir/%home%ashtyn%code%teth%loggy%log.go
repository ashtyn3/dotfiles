Vim�UnDo� 8��5+m"�4c���WX�X�T����x"I�1�                                     `K��    _�                             ����                                                                                                                                                                                                                                                                                                                                                             `K��     �                 �              5�_�                            ����                                                                                                                                                                                                                                                                                                                                                  V        `K��     �                import "fmt"       func main() {   	fmt.Println("vim-go")   }5�_�                           ����                                                                                                                                                                                                                                                                                                                                                  V        `K��     �                package main5�_�                           ����                                                                                                                                                                                                                                                                                                                                                  V        `K��     �                package 5�_�                            ����                                                                                                                                                                                                                                                                                                                                                  V        `K��     �                  �               �                  5�_�                            ����                                                                                                                                                                                                                                                                                                                                                  V        `K��     �                5�_�                           ����                                                                                                                                                                                                                                                                                                                                                  V        `K�     �               func Init()5�_�      	                     ����                                                                                                                                                                                                                                                                                                                                                  V        `K�     �                 defer logger.Sync()5�_�      
           	           ����                                                                                                                                                                                                                                                                                                                                                  V        `K�     �                package loggy �               "logger, err := zap.NewProduction()   if err != nil {   4  log.Fatalf("can't initialize zap logger: %v", err)�                 defer logger.Sync()   }5�_�   	              
           ����                                                                                                                                                                                                                                                                                                                                                  V        `K�    �               	   package loggy       func Init() {   #	logger, err := zap.NewProduction()   	if err != nil {   4		log.Fatalf("can't initialize zap logger: %v", err)   	}   	defer logger.Sync()   }5�_�   
                        ����                                                                                                                                                                                                                                                                                                                                                  V        `K�V     �                	defer logger.Sync()5�_�                           ����                                                                                                                                                                                                                                                                                                                                                  V        `K�c     �               4		log.Fatalf("can't initialize zap logger: %v", err)5�_�                           ����                                                                                                                                                                                                                                                                                                                                                  V        `K�h     �               6		log.Fatalf\n("can't initialize zap logger: %v", err)5�_�                           ����                                                                                                                                                                                                                                                                                                                                                  V        `K�h     �               5		log.Fatalfn("can't initialize zap logger: %v", err)5�_�                           ����                                                                                                                                                                                                                                                                                                                                                  V        `K�i     �               6		log.Fatalfln("can't initialize zap logger: %v", err)5�_�                            ����                                                                                                                                                                                                                                                                                                                                                  V        `K�m     �             5�_�                    
       ����                                                                                                                                                                                                                                                                                                                                                  V        `K�t     �   	            #	logger, err := zap.NewProduction()5�_�                    
       ����                                                                                                                                                                                                                                                                                                                                                  V        `K�v     �   	            	logger, err := zap.New()5�_�                    
        ����                                                                                                                                                                                                                                                                                                                                                  V        `K�    �                  package loggy       import (   	"log"       	"go.uber.org/zap"   )       func Init() {   $	logger, err := zap.NewDevelopment()   	if err != nil {   5		log.Fatalln("can't initialize zap logger: %v", err)   	}       }5�_�                           ����                                                                                                                                                                                                                                                                                                                                                  V        `K�     �             5�_�                            ����                                                                                                                                                                                                                                                                                                                                                  V        `K�     �                5�_�                    
       ����                                                                                                                                                                                                                                                                                                                                                  V        `K�    �                  package loggy       import (   	"log"       	"go.uber.org/zap"   )       func Init() {   $	logger, err := zap.NewDevelopment()   	if err != nil {   5		log.Fatalln("can't initialize zap logger: %v", err)   	}       	return logger   }5�_�                    	       ����                                                                                                                                                                                                                                                                                                                                                             `K�     �      
         func Init() {5�_�                    	       ����                                                                                                                                                                                                                                                                                                                                                             `K�     �      
         func Init()  {5�_�                    	       ����                                                                                                                                                                                                                                                                                                                                                             `K��     �      
         func Init() *zap.Logger  {5�_�                    	       ����                                                                                                                                                                                                                                                                                                                                                             `K��    �                  package loggy       import (   	"log"       	"go.uber.org/zap"   )       func Init() *zap.Logger {   $	logger, err := zap.NewDevelopment()   	if err != nil {   5		log.Fatalln("can't initialize zap logger: %v", err)   	}       	return logger   }5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             `K��     �               5		log.Fatalln("can't initialize zap logger: %v", err)5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             `K��     �               4		log.Fataln("can't initialize zap logger: %v", err)5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             `K��     �               3		log.Fatal("can't initialize zap logger: %v", err)5�_�                       -    ����                                                                                                                                                                                                                                                                                                                                                             `K��     �               4		log.Fatalf("can't initialize zap logger: %v", err)5�_�                        .    ����                                                                                                                                                                                                                                                                                                                                                             `K��    �                  package loggy       import (   	"log"       	"go.uber.org/zap"   )       func Init() *zap.Logger {   $	logger, err := zap.NewDevelopment()   	if err != nil {   6		log.Fatalf("can't initialize zap logger: %v\n", err)   	}       	return logger   }5��