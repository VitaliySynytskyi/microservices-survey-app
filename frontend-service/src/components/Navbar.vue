<template>
  <div class="navbar navbar-expand-lg fixed-top navbar-dark">
    <div class="container">
      <!-- Brand link to Home page with logo -->
      <router-link :to="{ name: 'Home' }" tag="a" class="navbar-brand">
        <span class="logo-text">Survey<span class="accent">App</span></span>
      </router-link>
      
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
        <ul class="navbar-nav ml-auto">
          <!-- Link to home -->
          <li class="nav-item">
            <router-link :to="{ name: 'Home' }" tag="a" class="nav-link">
              <i class="fas fa-home mr-1"></i> Головна
            </router-link>
          </li>
          
          <!-- Link to create a new survey -->
          <li class="nav-item">
            <router-link :to="{ name: 'SurveyForm' }" tag="a" class="nav-link">
              <i class="fas fa-plus-circle mr-1"></i> Створити
            </router-link>
          </li>
          
          <!-- Link to analytics dashboard -->
          <li class="nav-item">
            <router-link :to="{ name: 'Analytics' }" tag="a" class="nav-link">
              <i class="fas fa-chart-bar mr-1"></i> Аналітика
            </router-link>
          </li>
          
          <!-- User dropdown -->
          <li class="nav-item dropdown ml-2">
            <a class="nav-link dropdown-toggle user-dropdown" href="#" id="navbarDropdown" role="button" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
              <img src="https://ui-avatars.com/api/?name=User&background=4e73df&color=fff" alt="User" class="user-avatar mr-1">
              <span class="d-none d-md-inline-block">Користувач</span>
            </a>
            <div class="dropdown-menu dropdown-menu-right" aria-labelledby="navbarDropdown">
              <a class="dropdown-item" href="#">
                <i class="fas fa-user-circle mr-2"></i> Профіль
              </a>
              <a class="dropdown-item" href="#">
                <i class="fas fa-cog mr-2"></i> Налаштування
              </a>
              <div class="dropdown-divider"></div>
              <router-link :to="{ name: 'Auth' }" tag="a" class="dropdown-item">
                <i class="fas fa-sign-in-alt mr-2"></i> Увійти
              </router-link>
            </div>
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
.navbar {
  background-color: #3a3f58;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  padding: 15px 0;
}

.navbar-brand {
  font-size: 1.5rem;
  padding: 0;
}

.logo-text {
  font-weight: 700;
  letter-spacing: 0.5px;
}

.accent {
  color: #4ee0bc;
}

.nav-link {
  font-weight: 500;
  padding: 8px 15px;
  margin: 0 5px;
  transition: all 0.3s;
  border-radius: 5px;
}

.nav-link:hover {
  background-color: rgba(255, 255, 255, 0.1);
}

.create-link {
  background-color: #4e73df;
  color: white !important;
  border-radius: 20px;
  padding: 8px 20px;
}

.create-link:hover {
  background-color: #3a64d8;
  transform: translateY(-1px);
}

.user-dropdown {
  display: flex;
  align-items: center;
  padding: 5px 15px;
}

.user-avatar {
  width: 32px;
  height: 32px;
  border-radius: 50%;
  object-fit: cover;
  border: 2px solid rgba(255, 255, 255, 0.2);
}

.dropdown-menu {
  border: none;
  border-radius: 8px;
  box-shadow: 0 0.5rem 1rem rgba(0, 0, 0, 0.15);
  padding: 0.5rem 0;
  min-width: 14rem;
}

.dropdown-item {
  padding: 0.5rem 1.5rem;
  font-weight: 500;
}

.dropdown-item:hover {
  background-color: rgba(78, 115, 223, 0.1);
}

.dropdown-divider {
  margin: 0.5rem 0;
}

/* Add Font Awesome for icons */
@import url('https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.15.3/css/all.min.css');
</style>
