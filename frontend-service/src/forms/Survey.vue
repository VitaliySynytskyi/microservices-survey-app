<template>
  <div class="survey-form">
    <!-- Display error message if any -->
    <div v-if="errorMessage" class="alert alert-danger">{{ errorMessage }}</div>

    <!-- Survey name input -->
    <div class="form-group">
      <label class="col-form-label col-form-label-lg" for="name">Name</label>
      <input
        class="form-control form-control-lg"
        type="text"
        placeholder="Provide the name of the survey..."
        id="name"
        v-model="name"
      >
    </div>

    <!-- Questions input -->
    <div class="form-group">
      <label class="col-form-label col-form-label-lg" for="questions">Questions</label>
      <input
        v-for="(question, index) in questions"
        :key="index"
        class="form-control form-control-lg mb-3"
        type="text"
        placeholder="Enter a question..."
        :id="'question' + (index + 1)"
        v-model="questions[index]"
      >
      <small class="form-text text-muted">You must provide at least two questions.</small>
      <button class="btn btn-info btn-sm mt-3" @click="addQuestion">Add another question</button>
    </div>

    <!-- Create and cancel buttons -->
    <button type="button" class="btn btn-success btn-lg mt-3" @click="createSurvey">Create survey</button>
    <router-link :to="{ name: 'Home' }" tag="button" class="btn btn-link btn-sm mt-3 text-danger">Cancel</router-link>
  </div>
</template>

<script>
export default {
  name: 'SurveyForm',
  data() {
    return {
      name: "",
      questions: ["", ""],
      errorMessage: "",
    };
  },
  methods: {
    // Add an empty question input
    addQuestion() {
      this.questions.push("");
    },
    // Create a new survey
    createSurvey() {
      this.errorMessage = "";

      // Validate survey name
      if (this.name === "") {
        this.errorMessage = "The name is required.";
        return;
      }

      // Filter out empty questions
      const questions = this.questions.filter(question => question !== "").map(text => ({ text }));

      // Validate questions
      if (questions.length < 2) {
        this.errorMessage = "There must be at least two questions.";
        return;
      }

      const survey = { name: this.name, questions };

      // TODO: Make this URL configurable
      fetch('http://localhost:8081/surveys', {
        method: "POST",
        body: JSON.stringify(survey),
        headers: { "Content-type": "application/json" }
      })
        .then(res => Promise.all([res.status, res.json()]))
        .then(([status, data]) => {
          if (status !== 201) {
            throw new Error("Survey creation call resulted in result code: " + status);
          }
          this.$router.push({ name: 'Survey', params: { id: data.id }});
        })
        .catch(error => {
          this.errorMessage = "Unable to create survey. Please try again.";
          console.error(error);
        });
    }
  }
}
</script>
