govarsfromfiles
====

**Read files** with GO to get variable values.

For example from configuration.php with database variables configuration you can get host, database name, username and password.

##Usage example

You can have for example a file called configuration.php like this:

>configuration.php

    <?php
    $license = "this is the licence";
    $db_host = "localhost";
    $db_username = "this is username";
    $db_password = "this is the password";
    $db_name = "this is the database name";
    ?>


##How to get the package

    go get github.com/eduk2/govarsfromfiles

You can get the variable values with this code:

>main.go

    import (
        "fmt"
        v "github.com/eduk2/govarsfromfiles"
    )

    func main() {

        if v.File.SetPath("./configuration.php") {
            myValues := v.File.GetValues("$db_host", "$db_username", "$db_name", "$db_password")
            fmt.Println(myValues["$db_host"])
            fmt.Println(myValues["$db_username"])
            fmt.Println(myValues["$db_name"])
            fmt.Println(myValues["$db_password"])
        }

    }

If your variables use others types of quotes for example ' like this:


>configuration.php

    <?php
    $license = 'this is the licence';
    $db_host = 'localhost';
    $db_username = 'this is username';
    $db_password = 'this is the password';
    $db_name = 'this is the database name';
    ?>

You can configure how to get it with a regular expression in this way:

>main.go

    import (
        "fmt"
        v "github.com/eduk2/govarsfromfiles"
    )

    func main() {

        if v.File.SetPath("./configuration.php") {
            v.File.SetRegularExp(`'(.*)'`)
            myValues := v.File.GetValues("$db_host", "$db_username", "$db_name", "$db_password")
            fmt.Println(myValues["$db_host"])
            fmt.Println(myValues["$db_username"])
            fmt.Println(myValues["$db_name"])
            fmt.Println(myValues["$db_password"])
        }
    }
