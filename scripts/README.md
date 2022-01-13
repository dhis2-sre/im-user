# User scripts

All scripts rely on the environment variable $ACCESS_TOKEN being exported.

Valid access and refresh tokens can be retrieved by running `signIn.sh` (or `signInAdmin.sh`) and exported with `eval $(./signIn.sh)` (or `eval $(./signInAdmin.sh)`).

The following scripts require Admin access:
* [createGroup.sh](createGroup.sh)
* [addUserToGroup.sh](addUserToGroup.sh)
* [addClusterConfigToGroup.sh](addClusterConfigToGroup.sh)
