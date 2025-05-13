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
        v-model="survey.name"
      >
    </div>

    <!-- Survey description input -->
    <div class="form-group">
      <label class="col-form-label col-form-label-lg" for="description">Description</label>
      <textarea
        class="form-control"
        id="description"
        placeholder="Provide a description for the survey..."
        v-model="survey.description"
        rows="3"
      ></textarea>
    </div>

    <!-- Survey settings -->
    <div class="form-group">
      <div class="custom-control custom-switch mb-2">
        <input type="checkbox" class="custom-control-input" id="allowAnonymous" v-model="survey.allowAnonymous">
        <label class="custom-control-label" for="allowAnonymous">Allow anonymous responses</label>
      </div>
      
      <div class="form-group">
        <label for="expirationDate">Expiration Date (Optional)</label>
        <input type="date" class="form-control" id="expirationDate" v-model="expirationDate">
      </div>
      
      <div class="form-group">
        <label for="thankYouMessage">Thank You Message</label>
        <input type="text" class="form-control" id="thankYouMessage" v-model="survey.thankYouMessage" 
          placeholder="Message to show after completion">
      </div>
    </div>

    <!-- Questions section -->
    <h3 class="mt-4 mb-3">Questions</h3>
    
    <div v-for="(question, index) in survey.questions" :key="index" class="card mb-4">
      <div class="card-header d-flex justify-content-between align-items-center">
        <h5 class="mb-0">Question {{index + 1}}</h5>
        <button class="btn btn-sm btn-danger" @click="removeQuestion(index)" 
          v-if="survey.questions.length > 1">Remove</button>
      </div>
      
      <div class="card-body">
        <!-- Question text -->
        <div class="form-group">
          <label :for="'question-text-' + index">Question Text</label>
          <input
            class="form-control"
            type="text"
            :placeholder="'Enter question ' + (index + 1) + '...'"
            :id="'question-text-' + index"
            v-model="question.text"
          >
        </div>
        
        <!-- Question type -->
        <div class="form-group">
          <label :for="'question-type-' + index">Question Type</label>
          <select class="form-control" :id="'question-type-' + index" v-model="question.type">
            <option value="single_choice">Single Choice</option>
            <option value="multiple_choice">Multiple Choice</option>
            <option value="text">Text Answer</option>
            <option value="rating">Rating Scale</option>
            <option value="scale">Numeric Scale</option>
            <option value="date">Date</option>
          </select>
        </div>
        
        <!-- Required checkbox -->
        <div class="custom-control custom-switch mb-3">
          <input type="checkbox" class="custom-control-input" :id="'required-' + index" v-model="question.required">
          <label class="custom-control-label" :for="'required-' + index">Required question</label>
        </div>
        
        <!-- Question help text -->
        <div class="form-group">
          <label :for="'help-text-' + index">Help Text (Optional)</label>
          <input 
            type="text" 
            class="form-control" 
            :id="'help-text-' + index" 
            placeholder="Additional help text..." 
            v-model="question.helpText"
          >
        </div>
        
        <!-- Options for choice questions -->
        <div v-if="question.type === 'single_choice' || question.type === 'multiple_choice'" class="mt-3">
          <label>Options</label>
          <div v-for="(option, optIndex) in question.options" :key="'opt-' + index + '-' + optIndex" class="form-group">
            <div class="input-group mb-2">
              <input
                type="text"
                class="form-control"
                :placeholder="'Option ' + (optIndex + 1)"
                v-model="option.text"
              >
              <div class="input-group-append">
                <button class="btn btn-outline-danger" type="button" @click="removeOption(index, optIndex)" 
                  v-if="question.options.length > 2">×</button>
              </div>
            </div>
          </div>
          <button class="btn btn-sm btn-outline-primary" @click="addOption(index)">Add Option</button>
        </div>
        
        <!-- Scale configuration for rating and scale questions -->
        <div v-if="question.type === 'rating' || question.type === 'scale'" class="form-row mt-3">
          <div class="col">
            <label :for="'min-value-' + index">Minimum Value</label>
            <input type="number" class="form-control" :id="'min-value-' + index" v-model.number="question.minValue" min="0">
          </div>
          <div class="col">
            <label :for="'max-value-' + index">Maximum Value</label>
            <input type="number" class="form-control" :id="'max-value-' + index" v-model.number="question.maxValue" min="1">
          </div>
        </div>
        
        <!-- Text placeholder for text questions -->
        <div v-if="question.type === 'text'" class="form-group mt-3">
          <label :for="'placeholder-' + index">Placeholder Text</label>
          <input type="text" class="form-control" :id="'placeholder-' + index" v-model="question.placeholder"
            placeholder="Enter placeholder text...">
        </div>
      </div>
    </div>
    
    <div class="mb-4">
      <button class="btn btn-info" @click="addQuestion">Add another question</button>
      <small class="form-text text-muted mt-2">You must provide at least one question.</small>
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
      survey: {
        name: "",
        description: "",
        questions: [],
        allowAnonymous: true,
        thankYouMessage: "Thank you for completing the survey!"
      },
      expirationDate: "",
      errorMessage: "",
    };
  },
  created() {
    // Initialize with two empty questions
    this.addQuestion();
    this.addQuestion();
  },
  methods: {
    // Create a default question object
    createDefaultQuestion() {
      return {
        text: "",
        type: "single_choice",
        required: true,
        options: [
          { text: "Option 1" },
          { text: "Option 2" }
        ],
        minValue: 1,
        maxValue: 5,
        placeholder: "Enter your answer here...",
        helpText: ""
      };
    },
    // Add an empty question
    addQuestion() {
      this.survey.questions.push(this.createDefaultQuestion());
    },
    // Remove a question
    removeQuestion(index) {
      if (this.survey.questions.length > 1) {
        this.survey.questions.splice(index, 1);
      }
    },
    // Add an option to a choice question
    addOption(questionIndex) {
      const optionCount = this.survey.questions[questionIndex].options.length;
      this.survey.questions[questionIndex].options.push({
        text: `Option ${optionCount + 1}`
      });
    },
    // Remove an option from a choice question
    removeOption(questionIndex, optionIndex) {
      const options = this.survey.questions[questionIndex].options;
      if (options.length > 2) {
        options.splice(optionIndex, 1);
      }
    },
    // Create a new survey
    createSurvey() {
      this.errorMessage = "";

      // Validate survey name
      if (this.survey.name === "") {
        this.errorMessage = "The survey name is required.";
        return;
      }

      // Check for valid questions
      const validQuestions = this.survey.questions.filter(q => q.text.trim() !== "");
      if (validQuestions.length < 1) {
        this.errorMessage = "There must be at least one question with text.";
        return;
      }

      // Validate each question based on its type
      for (const question of this.survey.questions) {
        if (question.text.trim() === "") continue; // Skip empty questions
        
        if ((question.type === 'single_choice' || question.type === 'multiple_choice') && 
            question.options.some(opt => opt.text.trim() === "")) {
          this.errorMessage = "All options must have text.";
          return;
        }
        
        if ((question.type === 'rating' || question.type === 'scale') && 
            (question.minValue >= question.maxValue || !question.minValue || !question.maxValue)) {
          this.errorMessage = "Min value must be less than max value for rating/scale questions.";
          return;
        }
      }

      // Prepare the data for submission
      const surveyData = JSON.parse(JSON.stringify(this.survey));
      
      // Filter out empty questions
      surveyData.questions = validQuestions;
      
      // Add expiration date if set
      if (this.expirationDate) {
        const expiresAt = new Date(this.expirationDate);
        surveyData.expiresAt = Math.floor(expiresAt.getTime() / 1000);
      }

      // TODO: Make this URL configurable
      fetch('http://localhost:8081/surveys', {
        method: "POST",
        body: JSON.stringify(surveyData),
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

<style scoped>
.card-header {
  background-color: #f8f9fa;
  border-bottom: 1px solid #e9ecef;
  padding: 15px 20px;
}

.card {
  margin-bottom: 25px;
  border-radius: 10px;
  overflow: hidden;
  box-shadow: 0 3px 10px rgba(0, 0, 0, 0.05);
  border: none;
}

.card-body {
  padding: 20px;
}

.form-group {
  margin-bottom: 20px;
}

.form-control {
  transition: border-color 0.2s ease-in-out, box-shadow 0.2s ease-in-out;
}

.form-control:focus {
  border-color: #4e73df;
  box-shadow: 0 0 0 0.2rem rgba(78, 115, 223, 0.25);
}

.custom-control-input:checked ~ .custom-control-label::before {
  background-color: #4e73df;
  border-color: #4e73df;
}

.btn-info {
  background-color: #36b9cc;
  border-color: #36b9cc;
}

.btn-info:hover {
  background-color: #2ea7b9;
  border-color: #2ea7b9;
}

.btn-danger {
  background-color: #e74a3b;
  border-color: #e74a3b;
}

.btn-danger:hover {
  background-color: #e02d1b;
  border-color: #e02d1b;
}

.survey-form {
  background-color: #ffffff;
  border-radius: 10px;
  padding: 25px;
  box-shadow: 0 4px 15px rgba(0, 0, 0, 0.05);
  margin-bottom: 30px;
}

.input-group-append .btn {
  border-top-right-radius: 5px;
  border-bottom-right-radius: 5px;
}

.col-form-label {
  font-weight: 600;
}

.form-text {
  margin-top: 5px;
  color: #6c757d;
}

.alert {
  border-radius: 8px;
  padding: 15px 20px;
}
</style>
