<!-- eslint-disable -->
<template>
  <div class="landing-page">
    <!-- Hero section -->
    <section class="hero-section">
      <div class="container">
        <div class="row align-items-center">
          <div class="col-lg-6 hero-content">
            <h1 class="hero-title">Створюйте потужні<br><span class="accent-text">опитування</span> за хвилини</h1>
            <p class="hero-subtitle">Інноваційна платформа, яка трансформує спосіб збору та аналізу даних</p>
            <div class="hero-actions">
              <router-link :to="{ name: 'SurveyForm' }" class="btn btn-primary btn-lg mr-3">
                <i class="fas fa-rocket mr-2"></i>Почати створення
              </router-link>
              <router-link :to="{ name: 'Analytics' }" class="btn btn-outline-light btn-lg">
                <i class="fas fa-chart-line mr-2"></i>Демо аналітики
              </router-link>
            </div>
          </div>
          <div class="col-lg-6 d-none d-lg-block">
            <img src="https://images.unsplash.com/photo-1551288049-bebda4e38f71?ixlib=rb-1.2.1&auto=format&fit=crop&w=800&q=80" 
                 alt="Survey App Dashboard" class="img-fluid hero-image">
          </div>
        </div>
      </div>
      <div class="wave-divider">
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 1440 320">
          <path fill="#ffffff" fill-opacity="1" d="M0,128L48,133.3C96,139,192,149,288,154.7C384,160,480,160,576,138.7C672,117,768,75,864,74.7C960,75,1056,117,1152,122.7C1248,128,1344,96,1392,80L1440,64L1440,320L1392,320C1344,320,1248,320,1152,320C1056,320,960,320,864,320C768,320,672,320,576,320C480,320,384,320,288,320C192,320,96,320,48,320L0,320Z"></path>
        </svg>
      </div>
    </section>

    <!-- Features section -->
    <section class="features-section">
      <div class="container">
        <div class="section-header text-center mb-5">
          <h2 class="section-title">Розширені можливості</h2>
          <p class="section-subtitle">Повний набір інструментів для створення ідеальних опитувань</p>
        </div>
        
        <div class="row">
          <div class="col-md-4 mb-4">
            <div class="feature-card">
              <div class="feature-icon">
                <i class="fas fa-layer-group"></i>
              </div>
              <h3 class="feature-title">Різноманітні типи питань</h3>
              <p class="feature-description">Використовуйте одиночний вибір, множинний вибір, текстові поля, рейтингові шкали та багато іншого</p>
            </div>
          </div>
          
          <div class="col-md-4 mb-4">
            <div class="feature-card">
              <div class="feature-icon">
                <i class="fas fa-image"></i>
              </div>
              <h3 class="feature-title">Медіа вкладення</h3>
              <p class="feature-description">Додавайте зображення, відео та аудіо до питань для більшої наочності</p>
            </div>
          </div>
          
          <div class="col-md-4 mb-4">
            <div class="feature-card">
              <div class="feature-icon">
                <i class="fas fa-random"></i>
              </div>
              <h3 class="feature-title">Умовна логіка</h3>
              <p class="feature-description">Створюйте динамічні опитування з питаннями, що залежать від попередніх відповідей</p>
            </div>
          </div>
          
          <div class="col-md-4 mb-4">
            <div class="feature-card">
              <div class="feature-icon">
                <i class="fas fa-chart-bar"></i>
              </div>
              <h3 class="feature-title">Розширена аналітика</h3>
              <p class="feature-description">Отримуйте інтерактивні візуалізації та глибокі аналітичні звіти в реальному часі</p>
            </div>
          </div>
          
          <div class="col-md-4 mb-4">
            <div class="feature-card">
              <div class="feature-icon">
                <i class="fas fa-shield-alt"></i>
              </div>
              <h3 class="feature-title">Безпека та приватність</h3>
              <p class="feature-description">Збирайте анонімні відповіді та забезпечуйте захист персональних даних</p>
            </div>
          </div>
          
          <div class="col-md-4 mb-4">
            <div class="feature-card">
              <div class="feature-icon">
                <i class="fas fa-globe"></i>
              </div>
              <h3 class="feature-title">Багатомовна підтримка</h3>
              <p class="feature-description">Створюйте опитування різними мовами для глобальної аудиторії</p>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Use cases section -->
    <section class="use-cases-section">
      <div class="container">
        <div class="section-header text-center mb-5">
          <h2 class="section-title">Для чого використовувати</h2>
          <p class="section-subtitle">Наш сервіс підходить для різноманітних сценаріїв використання</p>
        </div>
        
        <div class="row">
          <div class="col-lg-3 col-md-6 mb-4">
            <div class="use-case-card">
              <div class="use-case-icon"><i class="fas fa-building"></i></div>
              <h4>Бізнес-дослідження</h4>
              <p>Оцінка задоволеності клієнтів, вивчення ринку та опитування співробітників</p>
            </div>
          </div>
          
          <div class="col-lg-3 col-md-6 mb-4">
            <div class="use-case-card">
              <div class="use-case-icon"><i class="fas fa-graduation-cap"></i></div>
              <h4>Освіта</h4>
              <p>Тестування знань студентів, оцінка курсів та збір відгуків</p>
            </div>
          </div>
          
          <div class="col-lg-3 col-md-6 mb-4">
            <div class="use-case-card">
              <div class="use-case-icon"><i class="fas fa-flask"></i></div>
              <h4>Наукові дослідження</h4>
              <p>Академічні опитування, польові дослідження та соціологічні експерименти</p>
            </div>
          </div>
          
          <div class="col-lg-3 col-md-6 mb-4">
            <div class="use-case-card">
              <div class="use-case-icon"><i class="fas fa-bullhorn"></i></div>
              <h4>Маркетинг</h4>
              <p>Збір відгуків про продукт, тестування концепцій та аналіз цільової аудиторії</p>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Coming soon section -->
    <section class="coming-soon-section">
      <div class="container">
        <div class="section-header text-center mb-5">
          <h2 class="section-title">Незабаром</h2>
          <p class="section-subtitle">Очікуйте на ці нові можливості у найближчому майбутньому</p>
        </div>
        
        <div class="row">
          <div class="col-md-4 mb-4">
            <div class="coming-soon-card">
              <div class="coming-soon-icon">
                <i class="fas fa-robot"></i>
              </div>
              <h3>AI-аналітика</h3>
              <p>Інтелектуальний аналіз відповідей, виявлення паттернів та прогнозування тенденцій</p>
              <span class="badge badge-pill badge-coming-soon">Скоро</span>
            </div>
          </div>
          
          <div class="col-md-4 mb-4">
            <div class="coming-soon-card">
              <div class="coming-soon-icon">
                <i class="fas fa-plug"></i>
              </div>
              <h3>Інтеграції</h3>
              <p>З'єднання з CRM, маркетинговими платформами, системами аналітики та іншими сервісами</p>
              <span class="badge badge-pill badge-coming-soon">Скоро</span>
            </div>
          </div>
          
          <div class="col-md-4 mb-4">
            <div class="coming-soon-card">
              <div class="coming-soon-icon">
                <i class="fas fa-money-bill-wave"></i>
              </div>
              <h3>Монетизація опитувань</h3>
              <p>Винагороджуйте респондентів та отримуйте дохід від проведення якісних опитувань</p>
              <span class="badge badge-pill badge-coming-soon">Скоро</span>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- CTA section -->
    <section class="cta-section">
      <div class="container">
        <div class="cta-content text-center">
          <h2>Готові перетворити спосіб збору даних?</h2>
          <p class="mb-4">Приєднуйтесь до тисячі компаній, які вже використовують нашу платформу для важливих рішень</p>
          <router-link :to="{ name: 'SurveyForm' }" class="btn btn-primary btn-lg">
            <i class="fas fa-paper-plane mr-2"></i>Створити перше опитування
          </router-link>
        </div>
      </div>
    </section>

    <!-- Survey list section -->
    <section class="survey-list-section" v-if="fetchedSurveys && filteredSurveys.length > 0">
      <div class="container">
        <div class="section-header d-flex justify-content-between align-items-center mb-4">
          <div>
            <h2 class="section-title">Доступні опитування</h2>
            <p class="section-subtitle">Перегляньте та пройдіть існуючі опитування</p>
          </div>
          
          <div class="d-flex">
            <div class="form-group mr-3 mb-0">
              <input type="text" class="form-control" placeholder="Пошук опитувань..." v-model="searchQuery">
            </div>
            <div class="form-group mb-0">
              <select class="form-control" v-model="sortOption">
                <option value="newest">Найновіші</option>
                <option value="oldest">Найстаріші</option>
                <option value="az">А-Я</option>
                <option value="za">Я-А</option>
              </select>
            </div>
          </div>
        </div>
        
        <div class="row">
          <div v-for="survey in filteredSurveys" :key="survey.id" class="col-lg-4 col-md-6 mb-4">
            <div class="survey-card">
              <div class="survey-status" :class="survey.expiresAt && new Date(survey.expiresAt * 1000) < new Date() ? 'expired' : 'active'">
                {{ survey.expiresAt && new Date(survey.expiresAt * 1000) < new Date() ? 'Закінчено' : 'Активно' }}
              </div>
              <h3 class="survey-title">{{ survey.name }}</h3>
              <p class="survey-description">{{ survey.description || 'Немає опису' }}</p>
              <div class="survey-meta">
                <span><i class="fas fa-question-circle"></i> {{ survey.questions.length }} питань</span>
                <span v-if="survey.expiresAt"><i class="fas fa-clock"></i> до {{ new Date(survey.expiresAt * 1000).toLocaleDateString() }}</span>
              </div>
              <router-link :to="{ name: 'Survey', params: { id: survey.id } }" class="btn btn-outline-primary btn-block mt-3">
                Переглянути опитування
              </router-link>
            </div>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
export default {
  name: 'Home',
  components: {},
  data() {
    return {
      surveys: [],
      errorMessage: "",
      fetchedSurveys: false,
      searchQuery: '',
      sortOption: 'newest',
      activeOnly: false
    };
  },
  computed: {
    filteredSurveys() {
      if (!this.surveys || !this.surveys.length) return [];
      
      // Filter surveys by search query and active status
      let result = [...this.surveys];
      
      if (this.searchQuery) {
        const query = this.searchQuery.toLowerCase();
        result = result.filter(survey => 
          survey.name.toLowerCase().includes(query) || 
          (survey.description && survey.description.toLowerCase().includes(query))
        );
      }
      
      if (this.activeOnly) {
        const now = Math.floor(Date.now() / 1000);
        result = result.filter(survey => 
          !survey.expiresAt || survey.expiresAt > now
        );
      }
      
      // Sort surveys
      switch (this.sortOption) {
        case 'newest':
          result.sort((a, b) => (b.createdAt || 0) - (a.createdAt || 0));
          break;
        case 'oldest':
          result.sort((a, b) => (a.createdAt || 0) - (b.createdAt || 0));
          break;
        case 'az':
          result.sort((a, b) => a.name.localeCompare(b.name));
          break;
        case 'za':
          result.sort((a, b) => b.name.localeCompare(a.name));
          break;
      }
      
      return result;
    }
  },
  created() {
    // Fetch surveys when the component is created
    // TODO: Make this URL configurable via an environment variable
    fetch("http://localhost:8081/surveys")
      .then(res => res.json())
      .then(data => {
        this.surveys = data;
        this.fetchedSurveys = true;
      })
      .catch(error => {
        this.errorMessage = "Не вдалося завантажити опитування!";
        // eslint-disable-next-line no-console
        console.error(error);
      });
  },
}
</script>

<style scoped>
/* Hero Section */
.hero-section {
  position: relative;
  background: linear-gradient(135deg, #5d4fbd 0%, #36b9cc 100%);
  color: white;
  padding: 120px 0 180px;
  overflow: hidden;
}

.hero-title {
  font-size: 3.2rem;
  font-weight: 800;
  line-height: 1.2;
  margin-bottom: 1.5rem;
}

.accent-text {
  color: #4ee0bc;
  position: relative;
  display: inline-block;
}

.accent-text:after {
  content: '';
  position: absolute;
  bottom: 5px;
  left: 0;
  width: 100%;
  height: 10px;
  background-color: rgba(78, 224, 188, 0.3);
  z-index: -1;
}

.hero-subtitle {
  font-size: 1.25rem;
  font-weight: 400;
  margin-bottom: 2.5rem;
  opacity: 0.9;
}

.hero-actions .btn {
  padding: 12px 30px;
  font-weight: 600;
  border-radius: 50px;
  transition: all 0.3s;
}

.hero-actions .btn-primary {
  background-color: #4ee0bc;
  border-color: #4ee0bc;
  color: #3a3f58;
}

.hero-actions .btn-primary:hover {
  background-color: #3cc9a7;
  border-color: #3cc9a7;
  transform: translateY(-3px);
  box-shadow: 0 10px 20px rgba(0, 0, 0, 0.1);
}

.hero-actions .btn-outline-light {
  border-width: 2px;
}

.hero-actions .btn-outline-light:hover {
  background-color: rgba(255, 255, 255, 0.1);
  transform: translateY(-3px);
  box-shadow: 0 10px 20px rgba(0, 0, 0, 0.1);
}

.hero-image {
  border-radius: 10px;
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.2);
  transform: perspective(1000px) rotateY(-5deg) rotateX(5deg);
  transition: all 0.5s;
}

.hero-image:hover {
  transform: perspective(1000px) rotateY(0deg) rotateX(0deg);
}

.wave-divider {
  position: absolute;
  bottom: 0;
  left: 0;
  width: 100%;
  line-height: 0;
}

/* Features Section */
.features-section {
  padding: 80px 0;
  background-color: #ffffff;
}

.section-header {
  margin-bottom: 60px;
}

.section-title {
  font-size: 2.5rem;
  font-weight: 700;
  margin-bottom: 15px;
  color: #3a3f58;
}

.section-subtitle {
  font-size: 1.25rem;
  color: #6c757d;
  max-width: 700px;
  margin: 0 auto;
}

.feature-card {
  background-color: #fff;
  border-radius: 15px;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.05);
  padding: 30px;
  height: 100%;
  transition: all 0.3s;
}

.feature-card:hover {
  transform: translateY(-10px);
  box-shadow: 0 15px 35px rgba(0, 0, 0, 0.1);
}

.feature-icon {
  width: 70px;
  height: 70px;
  background: linear-gradient(135deg, #5d4fbd 0%, #36b9cc 100%);
  border-radius: 20px;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 25px;
  font-size: 30px;
  color: white;
}

.feature-title {
  font-size: 1.25rem;
  font-weight: 600;
  margin-bottom: 15px;
  color: #3a3f58;
}

.feature-description {
  color: #6c757d;
  margin-bottom: 0;
}

/* Use Cases Section */
.use-cases-section {
  background-color: #f8f9fa;
  padding: 80px 0;
}

.use-case-card {
  background-color: #fff;
  border-radius: 10px;
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.05);
  padding: 25px;
  height: 100%;
  text-align: center;
  transition: all 0.3s;
}

.use-case-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 10px 25px rgba(0, 0, 0, 0.1);
}

.use-case-icon {
  font-size: 2.5rem;
  color: #5d4fbd;
  margin-bottom: 20px;
}

.use-case-card h4 {
  font-size: 1.25rem;
  font-weight: 600;
  margin-bottom: 15px;
  color: #3a3f58;
}

.use-case-card p {
  color: #6c757d;
  margin-bottom: 0;
}

/* Coming Soon Section */
.coming-soon-section {
  padding: 80px 0;
  background-color: #ffffff;
}

.coming-soon-card {
  background-color: #fff;
  border-radius: 15px;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.05);
  padding: 30px;
  height: 100%;
  position: relative;
  overflow: hidden;
  transition: all 0.3s;
  border: 2px dashed #e9ecef;
}

.coming-soon-card:hover {
  transform: translateY(-10px);
  box-shadow: 0 15px 35px rgba(0, 0, 0, 0.1);
  border-color: #5d4fbd;
}

.coming-soon-icon {
  width: 70px;
  height: 70px;
  background: linear-gradient(135deg, #ff7e5f 0%, #feb47b 100%);
  border-radius: 20px;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-bottom: 25px;
  font-size: 30px;
  color: white;
}

.coming-soon-card h3 {
  font-size: 1.25rem;
  font-weight: 600;
  margin-bottom: 15px;
  color: #3a3f58;
}

.coming-soon-card p {
  color: #6c757d;
  margin-bottom: 20px;
}

.badge-coming-soon {
  background-color: #ff7e5f;
  color: white;
  font-weight: 500;
  padding: 8px 15px;
  font-size: 0.75rem;
}

/* CTA Section */
.cta-section {
  padding: 80px 0;
  background: linear-gradient(135deg, #5d4fbd 0%, #36b9cc 100%);
  color: white;
}

.cta-content {
  max-width: 700px;
  margin: 0 auto;
}

.cta-content h2 {
  font-size: 2.5rem;
  font-weight: 700;
  margin-bottom: 20px;
}

.cta-content p {
  font-size: 1.25rem;
  opacity: 0.9;
  margin-bottom: 30px;
}

.cta-content .btn {
  padding: 12px 30px;
  font-weight: 600;
  border-radius: 50px;
  transition: all 0.3s;
  background-color: white;
  color: #5d4fbd;
  border-color: white;
}

.cta-content .btn:hover {
  transform: translateY(-3px);
  box-shadow: 0 10px 20px rgba(0, 0, 0, 0.2);
  background-color: white;
  color: #5d4fbd;
}

/* Survey List Section */
.survey-list-section {
  padding: 80px 0;
  background-color: #f8f9fa;
}

.survey-card {
  background-color: #fff;
  border-radius: 12px;
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.05);
  padding: 25px;
  height: 100%;
  position: relative;
  transition: all 0.3s;
}

.survey-card:hover {
  transform: translateY(-7px);
  box-shadow: 0 15px 30px rgba(0, 0, 0, 0.1);
}

.survey-status {
  position: absolute;
  top: 20px;
  right: 20px;
  padding: 5px 15px;
  border-radius: 50px;
  font-size: 0.75rem;
  font-weight: 600;
}

.survey-status.active {
  background-color: rgba(28, 200, 138, 0.1);
  color: #1cc88a;
}

.survey-status.expired {
  background-color: rgba(133, 135, 150, 0.1);
  color: #858796;
}

.survey-title {
  font-size: 1.35rem;
  font-weight: 600;
  margin-bottom: 12px;
  padding-right: 80px;
  color: #3a3f58;
  overflow: hidden;
  text-overflow: ellipsis;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
}

.survey-description {
  color: #6c757d;
  margin-bottom: 20px;
  overflow: hidden;
  text-overflow: ellipsis;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  height: 48px;
}

.survey-meta {
  display: flex;
  justify-content: space-between;
  margin-bottom: 10px;
  color: #858796;
  font-size: 0.875rem;
}

.survey-meta span {
  display: flex;
  align-items: center;
}

.survey-meta i {
  margin-right: 5px;
}

/* Responsive Media Queries */
@media (max-width: 991.98px) {
  .hero-title {
    font-size: 2.5rem;
  }
  
  .hero-section {
    padding: 80px 0 160px;
  }
  
  .section-title {
    font-size: 2rem;
  }
}

@media (max-width: 767.98px) {
  .hero-title {
    font-size: 2rem;
  }
  
  .hero-section {
    padding: 60px 0 140px;
    text-align: center;
  }
  
  .hero-actions {
    justify-content: center;
  }
  
  .section-title {
    font-size: 1.75rem;
  }
  
  .feature-card, .use-case-card, .coming-soon-card, .survey-card {
    padding: 20px;
  }
  
  .section-header {
    margin-bottom: 40px;
  }
}
</style>
