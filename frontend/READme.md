# *if you are using VS code... to view this file in a proper way press: ctrl+shift+V*

Hello

in this directory you will find the forntend of the server

---
this is our folder where ALL of our code goes
```
frontend
└───src
    ├───assets
    ├───components
    └───router
```

---

### package.json  &  package-lock.json 

 Don't mind about these files they are downloaded when you initialize a *vue* directory

 **DO NOT delete them because vue won't run without them**

---

### vite.config.js

This file is your Vite configuration file `vite.config.js` , and it plays a crucial role in how your Vue.js project is built and served

it imports required modules & Defines Vite Configuration

**DO NOT delete it because it's essential for your Vue project using Vite.**

---

**Modify it if:**
    
    * Your backend API runs on a different port or domain (change target).
    
    * You want to add more plugins or modify path aliases.

---

#### **1️⃣ `index.html` (Outside `src/`)**
This is the **entry point** of your Vue app. Vite uses this file as the main HTML template.


- it Defines a `<div id="app"></div>`, which is where Vue will mount your application.
- Loads `src/main.js`, which initializes Vue.

- 🔧 **Modify if needed:**
  - If you rename `main.js`, update the `<script>` tag accordingly.

---

#### **2️⃣ `main.js`**
This is the **main entry point** of your Vue application.

- it Imports `App.vue` (your main component).
- it Imports the `router` (to handle navigation).
- it Creates the Vue app and mounts it to `#app` in `index.html`.

---

#### **3️⃣ `App.vue`**
This is the **root component** of your Vue app.

- it Defines a simple welcome page with navigation links (`/login`, `/register`).
- Uses Vue Router’s `<router-link>` for navigation.
- Renders child components dynamically via `<router-view>`.

- 🔧 **Modify if needed:**
  - to include **more navigation links** as your project grows.

---

#### **4️⃣ `router/index.js`**
This file **configures and initializes Vue Router**, allowing navigation between different pages in your Vue app.

📌 **What It Does:**
- Imports the `Login.vue` and `Register.vue` components.
- Defines `routes` that map URL paths (`/login`, `/register`) to components.
- Creates and exports a **Vue Router instance** using `createWebHistory()` (which enables modern browser history mode).


- 🔧 **Modify if needed:**
  - Add **more routes** (e.g., `Home.vue`, `Dashboard.vue`).

---

#### **5️⃣ `components/` Directory (`Login.vue` & `Register.vue`)**
This folder contains **reusable UI components**, including authentication pages.


- **`Login.vue`**: this page only display what you write in the boxes & if you press the button redirects you to register page.
- **`Register.vue`**: contains a registration form for new users. and sends a backend request to save new users in the databse.

---
