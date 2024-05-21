<template>
  <div id="app">
    <!-- Navbar component for navigation -->
    <app-navbar />
    <!-- Main content area -->
    <div class="container">
      <router-view />
    </div>
  </div>
</template>

<script>
import Navbar from './components/Navbar.vue'

export default {
  name: 'App',
  components: {
    appNavbar: Navbar
  }
}
</script>

<style>
/* Global styles for the application can be added here */
</style>
