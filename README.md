# 4ID-test

To run the project you need to install ```docker``` and ```docker compose``` make sure you have ports ```8080 8081 6500 5432``` available.

After installing ```docker``` and ```docker compose```, to build and start the containers use the command 
```docker compose -p microservices up -d``` to stop and remove the containers use the command ```docker compose -p microservices down```. 

To access the resources, use ```http://localhost``` as basic URL, add the following depending of the action you want to do:
- Container ```micrologin```: to access the resource just add ```:8080/login``` or ```:8080/register```, usage will be described later.
- Container ```micrologout```: to access the resource use ```:8081/logout```, usage will be described later.
- To access database use ```localhost``` as your host, and the port ```6500```, username is ```root``` password ```123456```.


Services usage:
  - To use login service, use ```POST``` for the request, and send ```username:<your-username>, password:<your-password>``` in the body, if no user is found the first time it will register automatically
      with the credentials you send and login said user.
  - To register new users, use ```POST``` for the request and send ```username:<your-username>, password:<your-password>``` in the body.
  - To use logout service, use the ```POST```  for the request, and send ```username:<logged-username>``` in the body, if no user is found logged (already signed out) you will get a error.

On updates:
  After pulling new fixes or updates, stop the services and remove the containers, after that, delete the images of all and every service using docker GUI or the command
  ```docker image rm -f <container-id>```
