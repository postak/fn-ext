# Fn authentication extension 

This extension implement a simple and static authentication based on a middleware extension api.
It check for the presence of an **Authorization HTTP header** and check if the token correct with the respect of the one defined at startup via FN_TOKEN environment variable.

It secure apps by passing the correct Authorization HTTP header.

Without Authorization header

```
curl -H  http://127.0.0.1:8080/t/myapp/myfunc
{"error":{"message":"Invalid Authorization token."}}
```

With and Authorization header:

```
curl -H "Authorization: Bearer <mytoken>" http://127.0.0.1:8080/t/myapp/myfunc
OK
```
 
It works also with **fn** commands as it leverage the **fn context** feature.

Without a toekn or with a wrong one:

```
fn ls app

Fn: [GET /apps][401] ListApps default  &{Fields: Message:}

See 'fn <command> --help' for more information. Client version: 0.5.2
```

With correct token:

```
fn create ctx localfn --api-url http://127.0.0.1:8080
fn use ctx localfn
fn update ctx token mytoken
fn ls app
NAME
myapp
```

To build a fnserver image that includes this extension use the following command:

1. Clone this repo
2. Create a fnserver image using the following command
```
fn build-server -t imageuser/imagename:0.1
```
3. Push the image to you docker repo
```
docker push imageuser/imagename:0.1
```
4. Update the helm file *fn-helm/fn/values.yaml* :
* change pointer in the *fnserver* section for both *image* and *tag*. 
* add the environment *FN_AUTH* in the *env* section of *fnserver*
5. redeploy then new fn server using hrlm
        
You can find additional information [here](https://github.com/fnproject/fn-ext-example) and [here](https://github.com/fnproject/ext-auth)

