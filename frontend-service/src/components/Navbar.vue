<template>
  <div class="navbar navbar-expand-lg fixed-top navbar-dark bg-primary">
    <div class="container">
      <!-- Brand link to Home page -->
      <router-link :to="{ name: 'Home' }" tag="a" class="navbar-brand font-weight-bold">Surveys</router-link>
      <!-- Navbar toggler button for responsive design -->
      <button
        class="navbar-toggler"
        type="button"
        data-toggle="collapse"
        data-target="#navbarResponsive"
        aria-controls="navbarResponsive"
        aria-expanded="false"
        aria-label="Toggle navigation"
      >
        <span class="navbar-toggler-icon"></span>
      </button>
      <!-- Collapsible part of the navbar -->
      <div class="collapse navbar-collapse" id="navbarResponsive">
        <ul class="navbar-nav">
          <!-- Link to create a new survey -->
          <li class="nav-item">
            <router-link :to="{ name: 'SurveyForm' }" tag="a" class="nav-link">+ Create a new survey</router-link>
          </li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Navbar',
};
</script>

<style scoped>
/* Add any custom styles for the Navbar component here */
</style>
