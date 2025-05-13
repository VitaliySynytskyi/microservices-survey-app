<!-- eslint-disable -->
<template>
  <div class="analytics-dashboard">
    <div class="dashboard-header mb-4">
      <h1 class="display-5">Аналітична панель</h1>
      <p class="text-muted">Загальна статистика та аналіз опитувань</p>
    </div>

    <!-- Statistics cards -->
    <div class="row stats-cards mb-5">
      <div class="col-md-3">
        <div class="card stats-card">
          <div class="card-body">
            <div class="d-flex justify-content-between">
              <div>
                <h6 class="text-muted">Всього опитувань</h6>
                <h3 class="mb-0">{{ totalSurveys }}</h3>
              </div>
              <div class="stats-icon">
                <i class="fas fa-clipboard-list text-primary"></i>
              </div>
            </div>
            <div class="progress progress-sm mt-3">
              <div class="progress-bar bg-primary" role="progressbar" :style="{ width: '100%' }"></div>
            </div>
          </div>
        </div>
      </div>
      
      <div class="col-md-3">
        <div class="card stats-card">
          <div class="card-body">
            <div class="d-flex justify-content-between">
              <div>
                <h6 class="text-muted">Активні опитування</h6>
                <h3 class="mb-0">{{ activeSurveys }}</h3>
              </div>
              <div class="stats-icon">
                <i class="fas fa-check-circle text-success"></i>
              </div>
            </div>
            <div class="progress progress-sm mt-3">
              <div class="progress-bar bg-success" role="progressbar" :style="{ width: getPercentage(activeSurveys, totalSurveys) + '%' }"></div>
            </div>
          </div>
        </div>
      </div>
      
      <div class="col-md-3">
        <div class="card stats-card">
          <div class="card-body">
            <div class="d-flex justify-content-between">
              <div>
                <h6 class="text-muted">Відповідей всього</h6>
                <h3 class="mb-0">{{ totalResponses }}</h3>
              </div>
              <div class="stats-icon">
                <i class="fas fa-poll text-info"></i>
              </div>
            </div>
            <div class="progress progress-sm mt-3">
              <div class="progress-bar bg-info" role="progressbar" :style="{ width: '100%' }"></div>
            </div>
          </div>
        </div>
      </div>
      
      <div class="col-md-3">
        <div class="card stats-card">
          <div class="card-body">
            <div class="d-flex justify-content-between">
              <div>
                <h6 class="text-muted">Заверш. опитування</h6>
                <h3 class="mb-0">{{ completedSurveys }}</h3>
              </div>
              <div class="stats-icon">
                <i class="fas fa-calendar-check text-warning"></i>
              </div>
            </div>
            <div class="progress progress-sm mt-3">
              <div class="progress-bar bg-warning" role="progressbar" :style="{ width: getPercentage(completedSurveys, totalSurveys) + '%' }"></div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Charts and data -->
    <div class="row mb-4">
      <!-- Response Rate Chart -->
      <div class="col-lg-6">
        <div class="card mb-4">
          <div class="card-header">
            <h5 class="mb-0">Кількість відповідей за останні 7 днів</h5>
          </div>
          <div class="card-body">
            <div class="chart-container" style="height: 300px;">
              <canvas id="responsesChart" ref="responsesChart"></canvas>
            </div>
          </div>
        </div>
      </div>
      
      <!-- Survey Types Chart -->
      <div class="col-lg-6">
        <div class="card mb-4">
          <div class="card-header">
            <h5 class="mb-0">Типи питань в опитуваннях</h5>
          </div>
          <div class="card-body">
            <div class="chart-container" style="height: 300px;">
              <canvas id="questionTypesChart" ref="questionTypesChart"></canvas>
            </div>
          </div>
        </div>
      </div>
    </div>
    
    <!-- Recent surveys table -->
    <div class="card mb-4">
      <div class="card-header">
        <div class="d-flex justify-content-between align-items-center">
          <h5 class="mb-0">Останні активні опитування</h5>
          <router-link :to="{ name: 'Home' }" class="btn btn-sm btn-outline-primary">
            Переглянути всі
          </router-link>
        </div>
      </div>
      <div class="card-body p-0">
        <div class="table-responsive">
          <table class="table table-striped table-hover mb-0">
            <thead>
              <tr>
                <th>Назва</th>
                <th>Дата створення</th>
                <th>Питань</th>
                <th>Відповідей</th>
                <th>Статус</th>
                <th>Дії</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="survey in recentSurveys" :key="survey.id">
                <td>{{ survey.name }}</td>
                <td>{{ formatDate(survey.createdAt) }}</td>
                <td>{{ survey.questions.length }}</td>
                <td>{{ survey.responseCount || 0 }}</td>
                <td>
                  <span class="badge" :class="survey.expiresAt && new Date(survey.expiresAt * 1000) < new Date() ? 'badge-secondary' : 'badge-success'">
                    {{ survey.expiresAt && new Date(survey.expiresAt * 1000) < new Date() ? 'Закінчено' : 'Активно' }}
                  </span>
                </td>
                <td>
                  <router-link :to="{ name: 'Survey', params: { id: survey.id } }" class="btn btn-sm btn-outline-primary">
                    <i class="fas fa-eye"></i>
                  </router-link>
                </td>
              </tr>
              <tr v-if="!recentSurveys.length">
                <td colspan="6" class="text-center py-4">Немає доступних опитувань</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
// Import Chart.js if needed
// import Chart from 'chart.js';

export default {
  name: 'Analytics',
  data() {
    return {
      surveys: [],
      responses: [],
      loading: true,
      error: null,
      // Mock data for development
      totalSurveys: 0,
      activeSurveys: 0,
      completedSurveys: 0,
      totalResponses: 0,
      recentSurveys: [],
      responsesChartInstance: null,
      questionTypesChartInstance: null
    };
  },
  computed: {
    // Compute any derived metrics here
  },
  methods: {
    // Format a Unix timestamp to a readable date
    formatDate(timestamp) {
      if (!timestamp) return 'Невідомо';
      const date = new Date(timestamp * 1000);
      return date.toLocaleDateString('uk-UA');
    },
    
    // Calculate percentage
    getPercentage(value, total) {
      if (total === 0) return 0;
      return Math.round((value / total) * 100);
    },
    
    // Fetch surveys data
    fetchSurveys() {
      fetch("http://localhost:8081/surveys")
        .then(res => res.json())
        .then(data => {
          this.surveys = data;
          this.processSurveyData();
        })
        .catch(error => {
          this.error = "Не вдалося завантажити дані опитувань";
          // eslint-disable-next-line no-console
          console.error(error);
        })
        .finally(() => {
          this.loading = false;
        });
    },
    
    // Process survey data to extract statistics
    processSurveyData() {
      const now = Math.floor(Date.now() / 1000);
      
      this.totalSurveys = this.surveys.length;
      this.activeSurveys = this.surveys.filter(s => !s.expiresAt || s.expiresAt > now).length;
      this.completedSurveys = this.surveys.filter(s => s.expiresAt && s.expiresAt <= now).length;
      
      // Calculate total responses (mock data)
      this.totalResponses = this.surveys.reduce((sum, survey) => sum + (survey.responseCount || 0), 0);
      
      // Get recent surveys
      this.recentSurveys = [...this.surveys]
        .sort((a, b) => (b.createdAt || 0) - (a.createdAt || 0))
        .slice(0, 5);
      
      // Initialize charts after data is loaded
      this.$nextTick(() => {
        this.initCharts();
      });
    },
    
    // Initialize charts
    initCharts() {
      // Mock data for charts
      const last7Days = Array.from({length: 7}, (_, i) => {
        const date = new Date();
        date.setDate(date.getDate() - i);
        return date.toLocaleDateString('uk-UA', {weekday: 'short'});
      }).reverse();
      
      const responseData = [12, 19, 8, 15, 12, 25, 18];
      
      // Responses chart
      if (window.Chart && this.$refs.responsesChart) {
        if (this.responsesChartInstance) {
          this.responsesChartInstance.destroy();
        }
        
        this.responsesChartInstance = new window.Chart(this.$refs.responsesChart.getContext('2d'), {
          type: 'line',
          data: {
            labels: last7Days,
            datasets: [{
              label: 'Кількість відповідей',
              data: responseData,
              backgroundColor: 'rgba(78, 115, 223, 0.2)',
              borderColor: 'rgba(78, 115, 223, 1)',
              pointBackgroundColor: 'rgba(78, 115, 223, 1)',
              pointBorderColor: '#ffffff',
              pointRadius: 4,
              pointHoverRadius: 6,
              borderWidth: 2,
              tension: 0.3
            }]
          },
          options: {
            responsive: true,
            maintainAspectRatio: false
          }
        });
      }
      
      // Question types chart
      if (window.Chart && this.$refs.questionTypesChart) {
        if (this.questionTypesChartInstance) {
          this.questionTypesChartInstance.destroy();
        }
        
        const questionTypes = {
          'single_choice': 'Одиночний вибір',
          'multiple_choice': 'Множинний вибір',
          'text': 'Текстова відповідь',
          'rating': 'Рейтинг',
          'scale': 'Шкала',
          'date': 'Дата'
        };
        
        // Count question types across all surveys
        const typeCounts = {
          'single_choice': 45,
          'multiple_choice': 32,
          'text': 28,
          'rating': 15,
          'scale': 10,
          'date': 5
        };
        
        this.questionTypesChartInstance = new window.Chart(this.$refs.questionTypesChart.getContext('2d'), {
          type: 'doughnut',
          data: {
            labels: Object.values(questionTypes),
            datasets: [{
              data: Object.values(typeCounts),
              backgroundColor: [
                '#4e73df',
                '#1cc88a',
                '#36b9cc',
                '#f6c23e',
                '#e74a3b',
                '#6f42c1'
              ],
              borderWidth: 1
            }]
          },
          options: {
            responsive: true,
            maintainAspectRatio: false
          }
        });
      }
    }
  },
  created() {
    this.fetchSurveys();
  },
  mounted() {
    // Load Chart.js library
    if (!window.Chart) {
      const script = document.createElement('script');
      script.onload = () => this.initCharts();
      script.src = 'https://cdn.jsdelivr.net/npm/chart.js';
      document.head.appendChild(script);
    }
  }
}
</script>

<style scoped>
.analytics-dashboard {
  padding-bottom: 2rem;
}

.dashboard-header {
  margin-top: 1rem;
  margin-bottom: 2rem;
}

.stats-card {
  overflow: hidden;
  box-shadow: 0 0.15rem 1.75rem 0 rgba(58, 59, 69, 0.15);
  border: none;
  border-radius: 0.5rem;
  margin-bottom: 1.5rem;
  transition: transform 0.2s;
}

.stats-card:hover {
  transform: translateY(-5px);
}

.stats-icon {
  font-size: 2rem;
  opacity: 0.3;
}

.progress-sm {
  height: 0.5rem;
}

.chart-container {
  position: relative;
}

.table td, .table th {
  vertical-align: middle;
}

.table-hover tbody tr:hover {
  background-color: rgba(78, 115, 223, 0.05);
}

.badge {
  font-size: 0.75rem;
  padding: 0.4em 0.65em;
  font-weight: 600;
}
</style> 