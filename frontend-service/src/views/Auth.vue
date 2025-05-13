<template>
  <div class="auth-page">
    <div class="container">
      <div class="row justify-content-center">
        <div class="col-md-8 col-lg-6">
          <div class="card auth-card">
            <div class="card-header bg-transparent p-0">
              <!-- Tab navigation -->
              <ul class="nav nav-tabs nav-fill" role="tablist">
                <li class="nav-item">
                  <a class="nav-link" :class="{ active: activeTab === 'login' }" href="#" @click.prevent="activeTab = 'login'">
                    <i class="fas fa-sign-in-alt mr-2"></i>Вхід
                  </a>
                </li>
                <li class="nav-item">
                  <a class="nav-link" :class="{ active: activeTab === 'register' }" href="#" @click.prevent="activeTab = 'register'">
                    <i class="fas fa-user-plus mr-2"></i>Реєстрація
                  </a>
                </li>
              </ul>
            </div>
            
            <div class="card-body py-4 px-5">
              <!-- Login Form -->
              <div v-if="activeTab === 'login'">
                <h4 class="mb-4 text-center">Вхід в систему</h4>
                
                <div v-if="errorMessage" class="alert alert-danger">
                  {{ errorMessage }}
                </div>
                
                <form @submit.prevent="handleLogin">
                  <div class="form-group">
                    <label for="loginEmail">Електронна пошта</label>
                    <div class="input-group">
                      <div class="input-group-prepend">
                        <span class="input-group-text">
                          <i class="fas fa-envelope"></i>
                        </span>
                      </div>
                      <input 
                        type="email" 
                        class="form-control" 
                        id="loginEmail" 
                        v-model="loginForm.email" 
                        placeholder="Введіть email"
                        required
                      >
                    </div>
                  </div>
                  
                  <div class="form-group">
                    <label for="loginPassword">Пароль</label>
                    <div class="input-group">
                      <div class="input-group-prepend">
                        <span class="input-group-text">
                          <i class="fas fa-lock"></i>
                        </span>
                      </div>
                      <input 
                        :type="showPassword ? 'text' : 'password'" 
                        class="form-control" 
                        id="loginPassword" 
                        v-model="loginForm.password" 
                        placeholder="Введіть пароль"
                        required
                      >
                      <div class="input-group-append">
                        <button class="btn btn-outline-secondary" type="button" @click="showPassword = !showPassword">
                          <i :class="showPassword ? 'fas fa-eye-slash' : 'fas fa-eye'"></i>
                        </button>
                      </div>
                    </div>
                  </div>
                  
                  <div class="form-group form-check">
                    <input type="checkbox" class="form-check-input" id="rememberMe" v-model="loginForm.rememberMe">
                    <label class="form-check-label" for="rememberMe">Запам'ятати мене</label>
                    <a href="#" class="float-right">Забули пароль?</a>
                  </div>
                  
                  <button type="submit" class="btn btn-primary btn-block mt-4" :disabled="loading">
                    <span v-if="loading" class="spinner-border spinner-border-sm mr-2" role="status" aria-hidden="true"></span>
                    Увійти
                  </button>
                </form>
              </div>
              
              <!-- Register Form -->
              <div v-if="activeTab === 'register'">
                <h4 class="mb-4 text-center">Створення нового облікового запису</h4>
                
                <div v-if="errorMessage" class="alert alert-danger">
                  {{ errorMessage }}
                </div>
                
                <form @submit.prevent="handleRegister">
                  <div class="form-group">
                    <label for="registerName">Ім'я та прізвище</label>
                    <div class="input-group">
                      <div class="input-group-prepend">
                        <span class="input-group-text">
                          <i class="fas fa-user"></i>
                        </span>
                      </div>
                      <input 
                        type="text" 
                        class="form-control" 
                        id="registerName" 
                        v-model="registerForm.name" 
                        placeholder="Введіть повне ім'я"
                        required
                      >
                    </div>
                  </div>
                  
                  <div class="form-group">
                    <label for="registerEmail">Електронна пошта</label>
                    <div class="input-group">
                      <div class="input-group-prepend">
                        <span class="input-group-text">
                          <i class="fas fa-envelope"></i>
                        </span>
                      </div>
                      <input 
                        type="email" 
                        class="form-control" 
                        id="registerEmail" 
                        v-model="registerForm.email" 
                        placeholder="Введіть email"
                        required
                      >
                    </div>
                  </div>
                  
                  <div class="form-group">
                    <label for="registerPassword">Пароль</label>
                    <div class="input-group">
                      <div class="input-group-prepend">
                        <span class="input-group-text">
                          <i class="fas fa-lock"></i>
                        </span>
                      </div>
                      <input 
                        :type="showPassword ? 'text' : 'password'" 
                        class="form-control" 
                        id="registerPassword" 
                        v-model="registerForm.password" 
                        placeholder="Створіть пароль"
                        required
                      >
                      <div class="input-group-append">
                        <button class="btn btn-outline-secondary" type="button" @click="showPassword = !showPassword">
                          <i :class="showPassword ? 'fas fa-eye-slash' : 'fas fa-eye'"></i>
                        </button>
                      </div>
                    </div>
                    <small class="form-text text-muted">Пароль повинен містити щонайменше 8 символів, включаючи літери та цифри</small>
                  </div>
                  
                  <div class="form-group">
                    <label for="confirmPassword">Підтвердіть пароль</label>
                    <div class="input-group">
                      <div class="input-group-prepend">
                        <span class="input-group-text">
                          <i class="fas fa-lock"></i>
                        </span>
                      </div>
                      <input 
                        :type="showPassword ? 'text' : 'password'" 
                        class="form-control" 
                        id="confirmPassword" 
                        v-model="registerForm.confirmPassword" 
                        placeholder="Повторіть пароль"
                        required
                      >
                    </div>
                  </div>
                  
                  <div class="form-group form-check">
                    <input type="checkbox" class="form-check-input" id="termsCheck" v-model="registerForm.acceptTerms" required>
                    <label class="form-check-label" for="termsCheck">
                      Я погоджуюсь з <a href="#">Умовами використання</a> та <a href="#">Політикою конфіденційності</a>
                    </label>
                  </div>
                  
                  <button type="submit" class="btn btn-primary btn-block mt-4" :disabled="loading">
                    <span v-if="loading" class="spinner-border spinner-border-sm mr-2" role="status" aria-hidden="true"></span>
                    Зареєструватися
                  </button>
                </form>
              </div>
              
              <!-- Social login options -->
              <div class="social-login mt-4">
                <div class="text-center">
                  <p class="text-muted">Або увійдіть через соціальні мережі</p>
                  <div class="btn-group">
                    <button class="btn btn-outline-primary social-btn">
                      <i class="fab fa-facebook-f"></i>
                    </button>
                    <button class="btn btn-outline-danger social-btn">
                      <i class="fab fa-google"></i>
                    </button>
                    <button class="btn btn-outline-info social-btn">
                      <i class="fab fa-twitter"></i>
                    </button>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Auth',
  data() {
    return {
      activeTab: 'login',
      showPassword: false,
      loading: false,
      errorMessage: '',
      loginForm: {
        email: '',
        password: '',
        rememberMe: false
      },
      registerForm: {
        name: '',
        email: '',
        password: '',
        confirmPassword: '',
        acceptTerms: false
      }
    };
  },
  methods: {
    handleLogin() {
      this.loading = true;
      this.errorMessage = '';
      
      // Simulate API call
      setTimeout(() => {
        // Mock validation
        if (this.loginForm.email === 'admin@example.com' && this.loginForm.password === 'password') {
          // Success - redirect to home
          this.$router.push({ name: 'Home' });
        } else {
          // Error - show message
          this.errorMessage = 'Невірна електронна пошта або пароль';
        }
        this.loading = false;
      }, 1000);
    },
    
    handleRegister() {
      this.loading = true;
      this.errorMessage = '';
      
      // Validate confirm password
      if (this.registerForm.password !== this.registerForm.confirmPassword) {
        this.errorMessage = 'Паролі не співпадають';
        this.loading = false;
        return;
      }
      
      // Simulate API call
      setTimeout(() => {
        // Mock successful registration
        this.activeTab = 'login';
        // Reset form
        this.registerForm = {
          name: '',
          email: '',
          password: '',
          confirmPassword: '',
          acceptTerms: false
        };
        // Show success message
        alert('Реєстрація успішна! Тепер ви можете увійти.');
        this.loading = false;
      }, 1000);
    }
  }
}
</script>

<style scoped>
.auth-page {
  margin: 3rem 0;
}

.auth-card {
  border: none;
  border-radius: 10px;
  overflow: hidden;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.1);
}

.nav-tabs {
  border-bottom: none;
}

.nav-tabs .nav-link {
  border: none;
  font-weight: 500;
  padding: 15px;
  color: #6c757d;
  transition: all 0.3s;
}

.nav-tabs .nav-link.active {
  color: #4e73df;
  background-color: transparent;
  border-bottom: 3px solid #4e73df;
}

.input-group-text {
  background-color: transparent;
  border-right: none;
}

.form-control {
  border-left: none;
}

.input-group-text i {
  width: 16px;
  color: #6c757d;
}

.social-login {
  position: relative;
  margin-top: 30px;
  text-align: center;
}

.social-login:before {
  content: "";
  position: absolute;
  top: 50%;
  left: 0;
  right: 0;
  height: 1px;
  background-color: #e9ecef;
  z-index: 1;
}

.social-login p {
  display: inline-block;
  padding: 0 15px;
  background-color: white;
  position: relative;
  z-index: 2;
  margin-bottom: 15px;
}

.social-btn {
  width: 40px;
  height: 40px;
  padding: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 50%;
  margin: 0 5px;
  transition: all 0.3s;
}

.social-btn:hover {
  transform: translateY(-3px);
}
</style> 