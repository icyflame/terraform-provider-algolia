# terraform-provider-algolia

> Manage your Algolia configuration using terraform!
> Some code is borrowed from [bpicolo/terraform-provider-algolia][1]

Every Algolia index is a resource. Note that replicas are maintained in exactly the same way as
master indices. The internal master-replica relation is NOT mapped inside this provider.

## Importing existing Algolia resources

The import must be performed in two steps:

1. Update the `terraform state`

    ```sh
    $ terraform import algolia_index.local_reference "<index-name>"
    ```

    This updates the local terraform state.

2. Update the `.tf` files with the configurations of these resources

    After you have run the above command, you will find the file
    `state.<index-name>.timestamp` in the directory you ran this command.

    To get the terraform configuration, run:

    ```sh
    $ go run script/convert_to_conf.go state.<index-name>.timestamp
    ```

    This script will output the terraform configuration for this particular
    index. Copy this into the `.tf`file for the appropriate resource. Be sure to
    inspect the file to ensure that the types of all the variables is
    appropriate. (Arrays might not be printed with commas or the proper type)

[1]: https://github.com/bpicolo/terraform-provider-algolia
