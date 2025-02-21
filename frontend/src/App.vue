<template>
  <div class="welcome-container">
    <h1>Welcome to NTLink</h1>
    <nav>
      <router-link to="/login">Login</router-link> |
      <router-link to="/register">Register</router-link>
    </nav>
    <router-view class="router-view" />
   </div>
</template>

<script setup>
</script>
