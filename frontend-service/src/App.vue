<template>
  <div id="app">
    <!-- Navbar component for navigation -->
    <app-navbar />
    <!-- Main content area with transitions -->
    <div class="container">
      <transition name="fade" mode="out-in">
        <router-view />
      </transition>
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
/* Global styles for the application */
body {
  background-color: #f8f9fa;
  padding-top: 80px;
  color: #495057;
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
}

.container {
  max-width: 1140px;
  padding: 20px;
}

.card {
  border-radius: 8px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  border: none;
  margin-bottom: 20px;
  overflow: hidden;
}

.card-header {
  padding: 15px 20px;
  font-weight: 600;
}

.card-body {
  padding: 20px;
}

.btn {
  border-radius: 5px;
  padding: 8px 16px;
  font-weight: 500;
  transition: all 0.2s;
}

.btn-primary {
  background-color: #4e73df;
  border-color: #4e73df;
}

.btn-primary:hover, .btn-primary:focus {
  background-color: #3a64d8;
  border-color: #3a64d8;
}

.btn-success {
  background-color: #1cc88a;
  border-color: #1cc88a;
}

.btn-success:hover, .btn-success:focus {
  background-color: #18a87a;
  border-color: #18a87a;
}

.form-control {
  border-radius: 5px;
  padding: 10px 12px;
  height: auto;
}

/* Animate route transitions */
.fade-enter-active, .fade-leave-active {
  transition: opacity .3s;
}
.fade-enter, .fade-leave-to {
  opacity: 0;
}
</style>
