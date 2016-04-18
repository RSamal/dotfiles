Vim�UnDo� ��x$&�Grw�[G��)�<�#g��G�-�r��D   �                                  W�f    _�                     W   $    ����                                                                                                                                                                                                                                                                                                                                                             W�5     �   V   X   �      ,	Description("A user belonging to a member")5�_�                    W   $    ����                                                                                                                                                                                                                                                                                                                                                             W�6     �   V   X   �      &	Description("A user belonging to a ")5�_�                    W   )    ����                                                                                                                                                                                                                                                                                                                                                             W�9    �               �   package design       import (   $	. "github.com/goadesign/goa/design"   +	. "github.com/goadesign/goa/design/apidsl"   )       2// Authorize is the authorize resource media type.   Dvar Authorize = MediaType("application/vnd.authorize+json", func() {   ,	Description("Token authorization response")   	Attributes(func() {   3		Attribute("access_token", String, "access token")   C		Attribute("expires_in", Integer, "Time to expiration in seconds")   2		Attribute("token_type", String, "type of token")   !		Attribute("user", User, "user")   '		Attribute("tenant", Tenant, "tenant")       	})   	View("default", func() {   		Attribute("access_token")   		Attribute("expires_in")   		Attribute("token_type")   		Attribute("user")   		Attribute("tenant")   	})   })       *// Login is the Login resource media type.   <var Login = MediaType("application/vnd.login+json", func() {   	Description("")   	Attributes(func() {   %		Attribute("email", String, "email")   +		Attribute("password", String, "password")   	})   	View("default", func() {   		Attribute("email")   		Attribute("password")   	})   })       ,// Tenant is the tenant resource media type.   // GopherAcademy   >var Tenant = MediaType("application/vnd.tenant+json", func() {   '	Description("Tenant account in Congo")   	Attributes(func() {   #		Attribute("name", String, "name")   	})   	View("default", func() {   		Attribute("name")   	})   })       ,// Series is the series resource media type.   // GopherCon   >var Series = MediaType("application/vnd.series+json", func() {   =	Description("A series is a recurring event, e.g. GopherCon")   	Attributes(func() {   #		Attribute("name", String, "name")   	})   	View("default", func() {   		Attribute("name")   	})   })       *// Event is the event resource media type.   // GopherCon 2016   <var Event = MediaType("application/vnd.event+json", func() {   T	Description("An event is a specific instance of a conference, e.g. GopherCon 2016")   	Attributes(func() {   #		Attribute("name", String, "name")   +		Attribute("location", String, "location")   '		Attribute("url", String, "event URL")   1		Attribute("start_date", DateTime, "start_date")   -		Attribute("end_date", DateTime, "end_date")   	})   	View("default", func() {   		Attribute("name")   		Attribute("location")   		Attribute("url")   		Attribute("start_date")   		Attribute("end_date")   	})   })       (// User is the user resource media type.   :var User = MediaType("application/vnd.user+json", func() {   ,	Description("A user belonging to a tenant")   	Reference(UserPayload)   	Attributes(func() {   *		Attribute("id", Integer, "ID of record")   1		Attribute("href", String, "API href of record")   7		Attribute("first_name", String, "First name of user")   5		Attribute("last_name", String, "Last name of user")   3		Attribute("active", Boolean, "Activation status")   6		Attribute("validated", Boolean, "Validation status")   *		Attribute("role", String, "User's role")   !		Attribute("tenant_id", Integer)   >		Attribute("email", String, "Email address of user", func() {   			Format("email")   		})   	})       	View("default", func() {   		Attribute("id")   		Attribute("href")   		Attribute("first_name")   		Attribute("last_name")   		Attribute("email")   		Attribute("role")   		Attribute("tenant_id")   	})       	View("tenant", func() {   		Attribute("id")   		Attribute("href")   		Attribute("first_name")   		Attribute("last_name")   		Attribute("email")   		Attribute("active")   		Attribute("validated")   		Attribute("role")   		Attribute("tenant_id")   	})       	View("link", func() {   		Attribute("id")   		Attribute("href")   		Attribute("email")   	})   })5�_�                           ����                                                                                                                                                                                                                                                                                                                                                             W�     �      !   �      	Attributes(func() {5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             W�    �               �   package design       import (   $	. "github.com/goadesign/goa/design"   +	. "github.com/goadesign/goa/design/apidsl"   )       2// Authorize is the authorize resource media type.   Dvar Authorize = MediaType("application/vnd.authorize+json", func() {   ,	Description("Token authorization response")   	Attributes(func() {   3		Attribute("access_token", String, "access token")   C		Attribute("expires_in", Integer, "Time to expiration in seconds")   2		Attribute("token_type", String, "type of token")   !		Attribute("user", User, "user")   '		Attribute("tenant", Tenant, "tenant")       	})   	View("default", func() {   		Attribute("access_token")   		Attribute("expires_in")   		Attribute("token_type")   		Attribute("user")   		Attribute("tenant")   	})   })       *// Login is the Login resource media type.   <var Login = MediaType("application/vnd.login+json", func() {   	Description("")   	Attributes(func() {    		Attribute("id", Integer, "ID")   %		Attribute("email", String, "email")   +		Attribute("password", String, "password")   	})   	View("default", func() {   		Attribute("email")   		Attribute("password")   	})   })       ,// Tenant is the tenant resource media type.   // GopherAcademy   >var Tenant = MediaType("application/vnd.tenant+json", func() {   '	Description("Tenant account in Congo")   	Attributes(func() {   #		Attribute("name", String, "name")   	})   	View("default", func() {   		Attribute("name")   	})   })       ,// Series is the series resource media type.   // GopherCon   >var Series = MediaType("application/vnd.series+json", func() {   =	Description("A series is a recurring event, e.g. GopherCon")   	Attributes(func() {   #		Attribute("name", String, "name")   	})   	View("default", func() {   		Attribute("name")   	})   })       *// Event is the event resource media type.   // GopherCon 2016   <var Event = MediaType("application/vnd.event+json", func() {   T	Description("An event is a specific instance of a conference, e.g. GopherCon 2016")   	Attributes(func() {   #		Attribute("name", String, "name")   +		Attribute("location", String, "location")   '		Attribute("url", String, "event URL")   1		Attribute("start_date", DateTime, "start_date")   -		Attribute("end_date", DateTime, "end_date")   	})   	View("default", func() {   		Attribute("name")   		Attribute("location")   		Attribute("url")   		Attribute("start_date")   		Attribute("end_date")   	})   })       (// User is the user resource media type.   :var User = MediaType("application/vnd.user+json", func() {   ,	Description("A user belonging to a tenant")   	Reference(UserPayload)   	Attributes(func() {   *		Attribute("id", Integer, "ID of record")   1		Attribute("href", String, "API href of record")   7		Attribute("first_name", String, "First name of user")   5		Attribute("last_name", String, "Last name of user")   3		Attribute("active", Boolean, "Activation status")   6		Attribute("validated", Boolean, "Validation status")   *		Attribute("role", String, "User's role")   !		Attribute("tenant_id", Integer)   >		Attribute("email", String, "Email address of user", func() {   			Format("email")   		})   	})       	View("default", func() {   		Attribute("id")   		Attribute("href")   		Attribute("first_name")   		Attribute("last_name")   		Attribute("email")   		Attribute("role")   		Attribute("tenant_id")   	})       	View("tenant", func() {   		Attribute("id")   		Attribute("href")   		Attribute("first_name")   		Attribute("last_name")   		Attribute("email")   		Attribute("active")   		Attribute("validated")   		Attribute("role")   		Attribute("tenant_id")   	})       	View("link", func() {   		Attribute("id")   		Attribute("href")   		Attribute("email")   	})   })5�_�                    $       ����                                                                                                                                                                                                                                                                                                                                                             W�      �   #   &   �      	View("default", func() {5�_�                    %       ����                                                                                                                                                                                                                                                                                                                                                             W�%    �               �   package design       import (   $	. "github.com/goadesign/goa/design"   +	. "github.com/goadesign/goa/design/apidsl"   )       2// Authorize is the authorize resource media type.   Dvar Authorize = MediaType("application/vnd.authorize+json", func() {   ,	Description("Token authorization response")   	Attributes(func() {   3		Attribute("access_token", String, "access token")   C		Attribute("expires_in", Integer, "Time to expiration in seconds")   2		Attribute("token_type", String, "type of token")   !		Attribute("user", User, "user")   '		Attribute("tenant", Tenant, "tenant")       	})   	View("default", func() {   		Attribute("access_token")   		Attribute("expires_in")   		Attribute("token_type")   		Attribute("user")   		Attribute("tenant")   	})   })       *// Login is the Login resource media type.   <var Login = MediaType("application/vnd.login+json", func() {   	Description("")   	Attributes(func() {    		Attribute("id", Integer, "ID")   %		Attribute("email", String, "email")   +		Attribute("password", String, "password")   	})   	View("default", func() {   		Attribute("id")   		Attribute("email")   		Attribute("password")   	})   })       ,// Tenant is the tenant resource media type.   // GopherAcademy   >var Tenant = MediaType("application/vnd.tenant+json", func() {   '	Description("Tenant account in Congo")   	Attributes(func() {   #		Attribute("name", String, "name")   	})   	View("default", func() {   		Attribute("name")   	})   })       ,// Series is the series resource media type.   // GopherCon   >var Series = MediaType("application/vnd.series+json", func() {   =	Description("A series is a recurring event, e.g. GopherCon")   	Attributes(func() {   #		Attribute("name", String, "name")   	})   	View("default", func() {   		Attribute("name")   	})   })       *// Event is the event resource media type.   // GopherCon 2016   <var Event = MediaType("application/vnd.event+json", func() {   T	Description("An event is a specific instance of a conference, e.g. GopherCon 2016")   	Attributes(func() {   #		Attribute("name", String, "name")   +		Attribute("location", String, "location")   '		Attribute("url", String, "event URL")   1		Attribute("start_date", DateTime, "start_date")   -		Attribute("end_date", DateTime, "end_date")   	})   	View("default", func() {   		Attribute("name")   		Attribute("location")   		Attribute("url")   		Attribute("start_date")   		Attribute("end_date")   	})   })       (// User is the user resource media type.   :var User = MediaType("application/vnd.user+json", func() {   ,	Description("A user belonging to a tenant")   	Reference(UserPayload)   	Attributes(func() {   *		Attribute("id", Integer, "ID of record")   1		Attribute("href", String, "API href of record")   7		Attribute("first_name", String, "First name of user")   5		Attribute("last_name", String, "Last name of user")   3		Attribute("active", Boolean, "Activation status")   6		Attribute("validated", Boolean, "Validation status")   *		Attribute("role", String, "User's role")   !		Attribute("tenant_id", Integer)   >		Attribute("email", String, "Email address of user", func() {   			Format("email")   		})   	})       	View("default", func() {   		Attribute("id")   		Attribute("href")   		Attribute("first_name")   		Attribute("last_name")   		Attribute("email")   		Attribute("role")   		Attribute("tenant_id")   	})       	View("tenant", func() {   		Attribute("id")   		Attribute("href")   		Attribute("first_name")   		Attribute("last_name")   		Attribute("email")   		Attribute("active")   		Attribute("validated")   		Attribute("role")   		Attribute("tenant_id")   	})       	View("link", func() {   		Attribute("id")   		Attribute("href")   		Attribute("email")   	})   })5�_�      	              /       ����                                                                                                                                                                                                                                                                                                                                                             W�k     �   .   1   �      	Attributes(func() {5�_�      
           	   3       ����                                                                                                                                                                                                                                                                                                                                                             W�v     �   2   5   �      	View("default", func() {5�_�   	              
   4       ����                                                                                                                                                                                                                                                                                                                                                             W�z    �               �   package design       import (   $	. "github.com/goadesign/goa/design"   +	. "github.com/goadesign/goa/design/apidsl"   )       2// Authorize is the authorize resource media type.   Dvar Authorize = MediaType("application/vnd.authorize+json", func() {   ,	Description("Token authorization response")   	Attributes(func() {   3		Attribute("access_token", String, "access token")   C		Attribute("expires_in", Integer, "Time to expiration in seconds")   2		Attribute("token_type", String, "type of token")   !		Attribute("user", User, "user")   '		Attribute("tenant", Tenant, "tenant")       	})   	View("default", func() {   		Attribute("access_token")   		Attribute("expires_in")   		Attribute("token_type")   		Attribute("user")   		Attribute("tenant")   	})   })       *// Login is the Login resource media type.   <var Login = MediaType("application/vnd.login+json", func() {   	Description("")   	Attributes(func() {    		Attribute("id", Integer, "ID")   %		Attribute("email", String, "email")   +		Attribute("password", String, "password")   	})   	View("default", func() {   		Attribute("id")   		Attribute("email")   		Attribute("password")   	})   })       ,// Tenant is the tenant resource media type.   // GopherAcademy   >var Tenant = MediaType("application/vnd.tenant+json", func() {   '	Description("Tenant account in Congo")   	Attributes(func() {    		Attribute("id", Integer, "ID")   #		Attribute("name", String, "name")   	})   	View("default", func() {   		Attribute("id")   		Attribute("name")   	})   })       ,// Series is the series resource media type.   // GopherCon   >var Series = MediaType("application/vnd.series+json", func() {   =	Description("A series is a recurring event, e.g. GopherCon")   	Attributes(func() {   #		Attribute("name", String, "name")   	})   	View("default", func() {   		Attribute("name")   	})   })       *// Event is the event resource media type.   // GopherCon 2016   <var Event = MediaType("application/vnd.event+json", func() {   T	Description("An event is a specific instance of a conference, e.g. GopherCon 2016")   	Attributes(func() {   #		Attribute("name", String, "name")   +		Attribute("location", String, "location")   '		Attribute("url", String, "event URL")   1		Attribute("start_date", DateTime, "start_date")   -		Attribute("end_date", DateTime, "end_date")   	})   	View("default", func() {   		Attribute("name")   		Attribute("location")   		Attribute("url")   		Attribute("start_date")   		Attribute("end_date")   	})   })       (// User is the user resource media type.   :var User = MediaType("application/vnd.user+json", func() {   ,	Description("A user belonging to a tenant")   	Reference(UserPayload)   	Attributes(func() {   *		Attribute("id", Integer, "ID of record")   1		Attribute("href", String, "API href of record")   7		Attribute("first_name", String, "First name of user")   5		Attribute("last_name", String, "Last name of user")   3		Attribute("active", Boolean, "Activation status")   6		Attribute("validated", Boolean, "Validation status")   *		Attribute("role", String, "User's role")   !		Attribute("tenant_id", Integer)   >		Attribute("email", String, "Email address of user", func() {   			Format("email")   		})   	})       	View("default", func() {   		Attribute("id")   		Attribute("href")   		Attribute("first_name")   		Attribute("last_name")   		Attribute("email")   		Attribute("role")   		Attribute("tenant_id")   	})       	View("tenant", func() {   		Attribute("id")   		Attribute("href")   		Attribute("first_name")   		Attribute("last_name")   		Attribute("email")   		Attribute("active")   		Attribute("validated")   		Attribute("role")   		Attribute("tenant_id")   	})       	View("link", func() {   		Attribute("id")   		Attribute("href")   		Attribute("email")   	})   })5�_�   
                         ����                                                                                                                                                                                                                                                                                                                                                             W��     �                  		Attribute("id", Integer, "ID")5�_�                    $       ����                                                                                                                                                                                                                                                                                                                                                             W��     �   #   $          		Attribute("id")5�_�                     $        ����                                                                                                                                                                                                                                                                                                                                                             W�e    �               �   package design       import (   $	. "github.com/goadesign/goa/design"   +	. "github.com/goadesign/goa/design/apidsl"   )       2// Authorize is the authorize resource media type.   Dvar Authorize = MediaType("application/vnd.authorize+json", func() {   ,	Description("Token authorization response")   	Attributes(func() {   3		Attribute("access_token", String, "access token")   C		Attribute("expires_in", Integer, "Time to expiration in seconds")   2		Attribute("token_type", String, "type of token")   !		Attribute("user", User, "user")   '		Attribute("tenant", Tenant, "tenant")       	})   	View("default", func() {   		Attribute("access_token")   		Attribute("expires_in")   		Attribute("token_type")   		Attribute("user")   		Attribute("tenant")   	})   })       *// Login is the Login resource media type.   <var Login = MediaType("application/vnd.login+json", func() {   	Description("")   	Attributes(func() {   %		Attribute("email", String, "email")   +		Attribute("password", String, "password")   	})   	View("default", func() {   		Attribute("email")   		Attribute("password")   	})   })       ,// Tenant is the tenant resource media type.   // GopherAcademy   >var Tenant = MediaType("application/vnd.tenant+json", func() {   '	Description("Tenant account in Congo")   	Attributes(func() {    		Attribute("id", Integer, "ID")   #		Attribute("name", String, "name")   	})   	View("default", func() {   		Attribute("id")   		Attribute("name")   	})   })       ,// Series is the series resource media type.   // GopherCon   >var Series = MediaType("application/vnd.series+json", func() {   =	Description("A series is a recurring event, e.g. GopherCon")   	Attributes(func() {   #		Attribute("name", String, "name")   	})   	View("default", func() {   		Attribute("name")   	})   })       *// Event is the event resource media type.   // GopherCon 2016   <var Event = MediaType("application/vnd.event+json", func() {   T	Description("An event is a specific instance of a conference, e.g. GopherCon 2016")   	Attributes(func() {   #		Attribute("name", String, "name")   +		Attribute("location", String, "location")   '		Attribute("url", String, "event URL")   1		Attribute("start_date", DateTime, "start_date")   -		Attribute("end_date", DateTime, "end_date")   	})   	View("default", func() {   		Attribute("name")   		Attribute("location")   		Attribute("url")   		Attribute("start_date")   		Attribute("end_date")   	})   })       (// User is the user resource media type.   :var User = MediaType("application/vnd.user+json", func() {   ,	Description("A user belonging to a tenant")   	Reference(UserPayload)   	Attributes(func() {   *		Attribute("id", Integer, "ID of record")   1		Attribute("href", String, "API href of record")   7		Attribute("first_name", String, "First name of user")   5		Attribute("last_name", String, "Last name of user")   3		Attribute("active", Boolean, "Activation status")   6		Attribute("validated", Boolean, "Validation status")   *		Attribute("role", String, "User's role")   !		Attribute("tenant_id", Integer)   >		Attribute("email", String, "Email address of user", func() {   			Format("email")   		})   	})       	View("default", func() {   		Attribute("id")   		Attribute("href")   		Attribute("first_name")   		Attribute("last_name")   		Attribute("email")   		Attribute("role")   		Attribute("tenant_id")   	})       	View("tenant", func() {   		Attribute("id")   		Attribute("href")   		Attribute("first_name")   		Attribute("last_name")   		Attribute("email")   		Attribute("active")   		Attribute("validated")   		Attribute("role")   		Attribute("tenant_id")   	})       	View("link", func() {   		Attribute("id")   		Attribute("href")   		Attribute("email")   	})   })5��