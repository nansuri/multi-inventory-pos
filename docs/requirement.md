# Product Requirements
You are Experienced Software Engineer, i need you to create a web application which is basically inventory management for a restaurant, which able to track the stock for each ingredients based on the food prepared.

Unit Requirement:
* Tenant Management
  * This app will be able to serve multiple tenant
  * Each tenant can have it's own whole set of data such as
    * Super Employee (Owner)
    * Employee
    * Items (basically ingredients)
    * Recipe
    * Recipe Ingredients
    * Value per items
    * Currency used by the tenant
    * etc
  * User need to login first before can use this dashboard
* Stock/Inventory Management
  * User will be able to register items that will be under their inventory
  * User can register the item by using barcode / QR
  * Each inventory has:
    * price
    * stocks
    * etc
* Food & Recipe Management
  * User will be able to define their Product such as Food
  * Each food will have recipe that has relation with the Inventory that they have
  * User can define how many pax of each food will be prepared per day/week/month
  * These pax will also decrease the Inventory based on the ingredients
    * i.e. Fried Rice 10 pax, will need 1 KG of rice, 1 KG of onion, 500 Gr of Spices, etc
    * Once these already defined, inventory will also be affected
* Table Management & Order Management
  * Please define the requirements, basically as a restaurant they will have multiple tables, customer and orders, which will affect the recipe, and inventory
* Delivery Order 
  * Tenant can also track a delivery, etc
  * Please suggest requirements
* Main Dashboard
  * User will be able to track what items/inventory that is already run out
  * User can see their Money movement as well
* Other requirements, please suggest basically you can act as Product Manager

# Tech Stack Requirements
* Backend:
  * Go
  * Gorm, automigrate
  * Postgres
  * Redis
  * DDD Pattern
* Frontend
  * VueJS
  * Best Practice Design Pattern
* Deployment
  * Docker Container
  * 1 .env file for both BE and FE