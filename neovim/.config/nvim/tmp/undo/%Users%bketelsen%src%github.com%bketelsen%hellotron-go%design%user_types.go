Vim�UnDo� ��Z_~�)
]�P����Uف��p.X)l��Z�n                                     V�Q    _�                     
       ����                                                                                                                                                                                                                                                                                                                                                             V�D     �   	             	Attribute("fromEmail", func() {5�_�                    
       ����                                                                                                                                                                                                                                                                                                                                                             V�F     �   	            !	Attribute("from_Email", func() {5�_�                    
       ����                                                                                                                                                                                                                                                                                                                                                             V�G     �   	            !	Attribute("from_3mail", func() {5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             V�K     �               	Attribute("toEmail1", func() {5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             V�L     �                	Attribute("to_Email1", func() {5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             V�N     �               	Attribute("toEmail2", func() {5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             V�O     �                	Attribute("to_Email2", func() {5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             V�P    �                  package design       import (   +	. "github.com/goadesign/goa/design/apidsl"   )       R// IntroPayload defines the data structure used in the create bottle request body.   M// It is also the base type for the bottle media type used to render bottles.   0var IntroPayload = Type("IntroPayload", func() {   !	Attribute("from_email", func() {   		MinLength(10)   	})    	Attribute("to_email1", func() {   		MinLength(10)   	})    	Attribute("to_email2", func() {   		MinLength(10)   	})   })5��