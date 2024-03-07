## GVA Email Sending Plugin
#### Developer: Official GIN-VUE-ADMIN

### Usage Steps

#### 1. Go to the `initialize/router.go` file in the main GVA program. Add this plugin at the end of the `Routers` method according to your needs and security mode.
    Example:
    This plugin can use GVA's configuration file or directly hardcode content as configuration. It is recommended to add a configuration file structure for GVA and then pass the configuration.
	```go
	PluginInit(PrivateGroup, email.CreateEmailPlug(
		global.GVA_CONFIG.Email.To,
		global.GVA_CONFIG.Email.From,
		global.GVA_CONFIG.Email.Host,
		global.GVA_CONFIG.Email.Secret,
		global.GVA_CONFIG.Email.Nickname,
		global.GVA_CONFIG.Email.Port,
		global.GVA_CONFIG.Email.IsSSL,
	))
	```
    Similarly, you can hardcode when passing:
	```go
	PluginInit(PrivateGroup, email.CreateEmailPlug(
		"a@qq.com",
		"b@qq.com",
		"smtp.qq.com",
		"global.GVA_CONFIG.Email.Secret",
		"Login Key",
		465,
		true,
	))
	```

### 2. Configuration Explanation

#### 2-1 Global Configuration Struct
    // Where `From` and `Secret` are usually the username and password

    ```go
    type Email struct {
	    To       string  // Recipients: Multiple recipients separated by commas, e.g., a@qq.com b@qq.com. Use this as a parameter in formal development. This configuration is mainly used for sending error monitoring emails.
	    From     string  // Sender: Your own email address from which you want to send emails.
	    Host     string  // Server address: For example, smtp.qq.com. Please check the SMTP protocol for QQ or the email service you want to use.
	    Secret   string  // Secret: The key used for login. It's recommended not to use the email password. Apply for a key for login from the SMTP settings of your email provider.
	    Nickname string  // Nickname: Sender's nickname, usually your own email address. Customizable and optional.
	    Port     int     // Port: Please check the SMTP protocol for QQ or the email service you want to use. Often 465.
	    IsSSL    bool    // IsSSL: Whether to enable SSL.
    }
    ```

#### 2-2 Input Parameter Struct
    // Where `From` and `Secret` are usually the username and password

    ```go
    type Email struct {
        To      string `json:"to"`      // Recipient's email address
        Subject string `json:"subject"` // Email subject
        Body    string `json:"body"`    // Email content
    }
    ```

### 3. Method APIs

    - `utils.EmailTest(EmailTitle, EmailBody)` - Send a test email.
      Example: `utils.EmailTest("Test Email", "Test Email")`
      
    - `utils.ErrorToEmail(EmailTitle, EmailBody)` - Error monitoring.
      Example: `utils.ErrorToEmail("Test Email", "Test Email")`
      
    - `utils.Email(EmailTo, EmailTitle, EmailBody)` - Send a test email.
      Example: `utils.Email("a@qq.com,b@qq.com", "Test Email", "Test Email")`

### 4. Directly Accessible APIs

    - Test Endpoint: `/email/emailTest` [POST] - Swagger configured.

    - Email Sending Endpoint: `/email/emailSend` [POST] - Swagger configured.
      Input:
      ```go
      type Email struct {
          To      string `json:"to"`      // Recipient's email address
          Subject string `json:"subject"` // Email subject
          Body    string `json:"body"`    // Email content
      }
      ```
