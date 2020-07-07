# stolen-cars-management-system

This project stolen cars is build for the purpose of reporting stolen cars by car owners and every car is to be assigned to one police officer at a time and the police officer would resolve the status of the car accordingly that whether he Found the car or if he was unable to find the car after search.

Product Requirement:

1.  Car owners can report stolen cars. (Fulfilled)
2.  New stolen car cases should be automatically assigned to any free police officer. (Fulfilled)
3.  A police officer can only handle one stolen car case at a time. (Fulfilled)
4.  When the Police find a car, the case is marked as resolved and the responsible police officer become available to take a new stolen car case. (Fulfilled)
5.  The system should be able to assign unassigned stolen car cases automatically when a police officer becomes available. (Fulfilled)

Tech-Stack Used:

1. For Backend I'd used `Golang`.
2. Using `echo` framework of `Go`.
3. For Frontend I've chosen `ReactJS`.
4. Using `Material-UI`, popular React UI framework.
5. For Data-Stores, have used `MongoDB`.

For Running the project you should all the above tech-stacks installed.
For configuration the specific frontend and backend projects, please follow the README.md in the following folders, i.e. `stolencarsproject/client/consumer` and `stolencarsproject/server/`.

Once both the frontend and backend projects are setup. The worklow of the site will be,

Register(Sign-Up) yourself on-board using `Email` and `Name`, once registered then you can also `Sign-In` using only `Email`, then you can add stolen cars with details as car owners and can monitor once your cars is added and if any available Police Officer will be there then your stolen car will be assigned to him when you submit the stolen car details. Similarly, the Police officer can also `Sign-Up` and `Sign-In` as well and if any unassigned car will be available then will be assigned to them one car at a time and then after search of the they can resolve the stolen car status to `Found` or `Not-Found` accordingly.
