# Resource User tesing

<https://www.terraform.io/docs/extend/testing/acceptance-tests/index.html>

#### BASIC COMMANDS TO RUN TEST FILE

1-go test <br/>
2-go test -cover (to get idea about how much percentage of your code is tested) <br />
3-go test ./... (to run all the test file in a particular folder) <br />

//In our case we are doing acceptance testing <br />

### <strong> STEPS</strong>

1-make TF_ACC = true (set environment variable,this is to run the acceptance testing). <br />

2-Hashicorp has provider some inbuilt packages which we can use to implemet our testing ie. resource("github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource") <br />

3-we set up a resource.Test and provide it with the following: <br />
PreCheck,Providers,CheckDestroy,Steps. <br />

4- In each steps, we can provide couple of things <br />
Config,Check. <br />

### DIFFERENT TESTING FUNCTION


<strong>1. TestAccItem_Basic </strong>

- It calls the destroy and exist function to test delete and create function <br />
-  On providing wrong data in TestAccItem_Basic(CREATE,READ,DELETE), test must fails.

<strong>2. TestAccItem_Update</strong>

- It is the testing function for update <br />
- On passing incorrect data wrt  existing user in TestAccItem_Update(UPDATE), test must fails.