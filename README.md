# Test 1 for Flugel.it

This repository holds the files needed to run the **Test1** requested by Flugel.it to Marco de Martinis.

The requirements are as follows:

- Create Terraform code to create an S3 bucket and an EC2 instance. Both resources must be tagged with Name=Flugel, Owner=InfraTeam.
- Using Terratest, create the test automation for the Terraform code, validating that both resources are tagged properly.
- Setup Github Actions to run a pipeline to validate this code.
- Publish your code in a public GitHub repository, and share a Pull Request with your code. Do not merge into master until the PR is approved.
- Include documentation describing the steps to run and test the automation.

## How it Works

The solution is composed by Terraform, Go, and YML files, in order to deploy the infrastructure, test it and automate the entire process using GitHub Actions.

As you may notice, the branch *main* is empty. This is because all files were committed to the branch *pipeline* and are pending to merge with a Pull Request.

The GitHub Action workflow for CI runs on Pull Request creation. As first step, it lints the entire code base, to verify that everything has been written following the highest standards. Then it proceeds to prepare the runner to execute the test with Terratest.

---

### Running the solution locally

#### Requirements and Dependencies

To run the code locally, without any change, you must have a set of AWS credentials that are allowed to read and write the AWS S3 Bucket "terraform-state-mdemartinis", otherwise you will need to modifiy [backend.tf] to point to a new bucket or delete the ***backend*** definition to save the Terraform state locally.

It is required to have Terraform and Go installed, with minimum versions as follows (based on versions used to develop the solution):
- Terraform: 1.0.7
- Go: 1.17.1

Additionally, if you want to clone this repository, you will need:
- Git: 2.33

#### Clone repository

You can clone this repository and select the branch *pipeline* with the following commands:
```bash
git clone https://github.com/mdemartinis/GithubActions.git
cd ./GithubActions/
git checkout pipeline
```

To execute Terraform commands, you will want to move to the Terraform folder first, and then you will be able to initialize it using:
```bash
terraform init
```

After initializing Terraform, you will now be able to run any of the following commands:
```bash
terraform validate
terraform plan
terraform apply
terraform destroy
```

If you want to run the full test, you have to move first to the Terratest folder, initialize `go mod` and download all the dependencies. Then, execute the test. Use the following commands:
```bash
cd ./terratest/
go mod init <MODULE_NAME>
go mod tidy
go test
```

Replace `<MODULE_NAME>` with your desired module name, tipically, your repository name in the form of: **github.com/mdemartinis/GithubActions**

Always remember to verify that the run has completed successfully and the cloud resources have been destroyed, to not incur in undesired billing.


[//]: #

   [backend.tf]: <https://github.com/mdemartinis/GithubActions/blob/pipeline/terraform/backend.tf>
