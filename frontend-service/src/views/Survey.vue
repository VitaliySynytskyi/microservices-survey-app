<template>
  <div class="home">
    <!-- Display error message if any -->
    <div v-if="errorMessage" class="alert alert-danger">{{ errorMessage }}</div>
    <!-- Render the survey using the Survey component if survey data is available -->
    <app-survey v-if="survey" :survey="survey"></app-survey>
  </div>
</template>

<script>
import Survey from '../components/Survey.vue';

export default {
  name: 'Survey',
  components: {
    'app-survey': Survey,
  },
  data() {
    return {
      survey: null,
      errorMessage: "",
    };
  },
  created() {
    // Fetch survey data when the component is created
    // TODO: Make this URL configurable via an environment variable
    fetch(`http://localhost:8081/surveys/${this.$route.params.id}`)
      .then(res => Promise.all([res.status, res.json()]))
      .then(([status, data]) => {
        if (status === 404) {
          this.errorMessage = "Survey not found.";
          return;
        }
        this.survey = data;
      })
      .catch(error => {
        this.errorMessage = "Unable to load survey!";
        console.error(error);
      });
  },
}
</script>
