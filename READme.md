# *if you are using VS code... to view this file in a proper way press: ctrl+shift+V*

Hello

in this directory you will find the backend of the server

---

these are our project files & folders

```

├───backend
│   ├───db
│   │   ├───migrations
│   │   │   └───sqlite
│   │   └───sqlite
│   └───pkg
│       ├───database
│       ├───handlers
│       └───structs
└───frontend
    └───src
        ├───assets
        ├───components
        └───router
```

Each Directory contains a Readme file. PLEASE CHECK THEM OUT    
---

## to run this simple app

--> open 2 terminals and navigate to our folder `social-network-demo`

---
in the first terminal write the below commands

these commands will run the backend of the app

```
cd .\backend\
go run server.go
```

---

in the second terminal we will run the frontend

please be sure you have installed VUE@3

```
cd .\frontend\
npm install 
npm run dev
```
---

then navigate to `http://localhost:5173` & enjoy