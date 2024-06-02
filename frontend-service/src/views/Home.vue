<template>
  <div class="home">
    <!-- Display error message if any -->
    <div v-if="errorMessage" class="alert alert-danger">{{ errorMessage }}</div>
    <!-- Display a warning message if no surveys are available after fetching -->
    <div v-if="fetchedSurveys && !surveys.length" class="alert alert-warning">There are currently no surveys.</div>
    <!-- Render each survey using the Survey component -->
    <app-survey v-for="survey in surveys" :survey="survey" :key="survey.id"></app-survey>
  </div>
</template>

<script>
import Survey from '../components/Survey.vue';

export default {
  name: 'Home',
  components: {
    'app-survey': Survey,
  },
  data() {
    return {
      surveys: [],
      errorMessage: "",
      fetchedSurveys: false,
    };
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
        this.errorMessage = "Unable to load surveys!";
        console.error(error);
      });
  },
}
</script>
