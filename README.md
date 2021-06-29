This Terraform provider enables create, read, update, delete, and import operations for Microspft office365 users.

## Requirements

- [Go](https://golang.org/doc/install) >= 1.16 (To build the provider plugin)
- [Terraform](https://www.terraform.io/downloads.html) >= 0.13.x
- Application : [Micorsft office365](https://www.office.com/)
- [Office365 API documentation](https://docs.microsoft.com/en-us/graph/overview)

## Application account

#### Setup

1. Create Mirosoft office365 account
1. Go to the Join the [Microsoft 365 Developer Program](https://developer.microsoft.com/en-us/microsoft-365/dev-program) page and join the developer Program
2. Set up a [Microsoft 365 developer sandbox subscription](https://docs.microsoft.com/en-us/office/developer-program/microsoft-365-developer-program-get-started)
3. Register your application in AzureAd
   - Use the Azure Management Portal to register your application in Azure AD
   - For reference documentation click [here](https://docs.microsoft.com/en-us/office/office-365-management-api/get-started-with-office-365-management-apis#prerequisites)

#### API Authentication
1. To authenticate API, we need these credentials: ObjectID ,TenantID and ClientSecret.
2.  Generate a new key for your application by following [this](https://docs.microsoft.com/en-us/office/office-365-management-api/get-started-with-office-365-management-apis#generate-a-new-key-for-your-application)


## Building The Provider 
1. Clone the repository, add all the dependencies and create a vendor directory that contains all dependencies. For this, run the following commands: <br>
```
cd terraform-provider-office365
go mod init terraform-provider-office365
go mod tidy
go mod vendor
```

## Managing terraform plugins
*For Windows:*
1. Run the following command to create a vendor sub-directory (`%APPDATA%/terraform.d/plugins/${host_name}/${namespace}/${type}/${version}/${OS_ARCH}`) which will consist of all terraform plugins. <br> 
Command: 
```bash
mkdir -p %APPDATA%/terraform.d/plugins/office365.com/users/office365/1.0.0/windows_amd64
```
2. Run `go build -o terraform-provider-office365.exe` to generate the binary in present working directory. <br>
3. Run this command to move this binary file to the appropriate location.
 ```
 move terraform-provider-office365.exe %APPDATA%\terraform.d\plugins\office356.com\users\office365\1.0.0\windows_amd64
 ``` 
<p align="center">[OR]</p>
 
3. Manually move the file from current directory to destination directory (`%APPDATA%\terraform.d\plugins\office365.com\users\office365\1.0.0\windows_amd64`).<br>

## Working with terraform

### Application Credential Integration in terraform
1. Add `terraform` block and `provider` block as shown in [example usage](#example-usage).
2. Get credentials: ObjectID ,TenantID and ClientSecret. For this, visit https://docs.microsoft.com/en-us/azure/active-directory/develop/quickstart-register-app#add-credentials
3. Assign the above credentials to the respective field in the `provider` block.

### Basic Terraform Commands
1. `terraform init` - To initialize a working directory containing Terraform configuration files.
2. `terraform plan` - To create an execution plan. Displays the changes to be done.
3. `terraform apply` - To execute the actions proposed in a Terraform plan. Apply the changes.

### Create User
1. Add the `display_name`, `password`, `mail_nick_name`, `user_principal_name`,  in the respective field in `resource` block as shown in [example usage](#example-usage).
2. Run the basic terraform commands.<br>
3. On successful execution, Created new office365 user

### Update the user
1. Update the data of the user in the `resource` block as show in [example usage](#example-usage) and run the basic terraform commands to update user. 
3. User is not allowed to update `password` and `user_principal_name`.
2. To activate user, set account_enabled=true and vice-versa

### Read the User Data
Add `data` and `output` blocks as shown in the [example usage](#example-usage) and run the basic terraform commands.

### Delete the user
Delete the `resource` block of the user and run `terraform apply`.

### Import a User Data
1. Write manually a `resource` configuration block for the user as shown in [example usage](#example-usage). Imported user will be mapped to this block.
2. Run the command `terraform import office365_user.exmaple  [user_principal_name]` to import user.
3. Run `terraform plan`, if output shows `0 to addd, 0 to change and 0 to destroy` user import is successful.

## Example Usage<a id="example-usage"></a>
```
terraform {
  required_providers {
    office365 = {
      version = "~> 0.1"
      source  = "Raghvendra/clevertap/office365"
    }
  }
}

provider "office365" {
   client_id     = "_Replace_office365_Client_ID_"
   client_secret = "_Replace_office365_Client_Secret_"
   tenant_id     ="_Replace_office365_Tenant_ID_"

}

#Creating User
resource "office365_user_manage" "example" {
   display_name        ="user full name"
   mail_nick_name      ="nick name"
   user_principal_name ="example@<officce365domain>.onmicrosoft.com"
   password            ="********"
   account_enabled     ="true"
}


#Get user Information
data "office365_users" "example" {
    userprincipalname ="example@<officce365domain>.onmicrosoft.com"
}
 output "example" {
      value          = data.office365_users.example
}

```

### Argument Reference
- ``display_name`` (required,String) The name displayed in the address book for the user. This is usually the combination of the user's first name, middle initial and last name.<br/>

- ``user_principal_name`` (required,String)The user principal name (UPN) of the user. The UPN is an Internet-style login name for the user based on the Internet standard RFC 822. By convention, this should map to the user's email name. The general format is alias@domain, where domain must be present in the tenant's collection of verified domains. This property is required when a user is created. 

- ``password`` (required,String) password for created user.<br/>
- ``force_change_password_nextsignin`` (required,Boolean) It will force to change the password when user singin for first time.

- ``account_enabled`` (optional,Boolean) Takes only true/false as input and allow us to activate/deactivate account(Default:true).

- ``given_name`` (optional,String) first name of user.

- ``surname`` (optional,String) second name of user.

- ``jobtitle`` (optional,String) The user's job title. Maximum length is 128 characters.

- ``office_location``(optional,String) office location of user.

- ``mail_nick_name`` (optional,String) The mail alias for the user.

- ``mobile_phone`` (optional,String) The primary cellular telephone number for the user.
- ``postal_code`` (optional,String)The postal code for the user's postal address. The postal code is specific to the user's country/region.

- ``state`` (optional,String)The state or province in the user's address.
- ``preferred_language`` (optional,String)The preferred language for the user. Should follow ISO 639-1 Code; for example "en-US".
- ``street_address`` (optional,String)The street address of the user's place of business.
- ``usage_location`` (optional,String)A two letter country code (ISO standard 3166). Required for users that will be assigned licenses due to legal requirement to check for availability of services in countries. Examples include: "US", "JP", and "GB".


### Exceptions
- In import, only specific arguments are read by API. so to manage any arguments other than these, we have to enter manually.

-  Refere [this](https://docs.microsoft.com/en-us/graph/api/user-get?view=graph-rest-1.0&tabs=http#response-1) for arguments list.