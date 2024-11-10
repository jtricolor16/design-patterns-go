package main

import "fmt"

func main() {
	fmt.Println("builder")
	builderInstance := NewRunRelease()
	directorDeploymentExec := NewDirectorDeployment(builderInstance)
	directorDeploymentExec.ManageBuilder()
}
