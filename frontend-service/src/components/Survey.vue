<template>
  <div class="survey">
    <div class="card mb-4">
      <div class="card-header bg-primary text-white">
        <h4>{{ survey.name }}</h4>
        <p v-if="survey.description" class="mb-0">{{ survey.description }}</p>
      </div>
      <div class="card-body">
        <div v-if="errorMessage" class="alert alert-danger">{{ errorMessage }}</div>
        <div v-if="statusMessage" class="alert alert-success">{{ statusMessage }}</div>
        
        <template v-if="showQuestions">
          <!-- eslint-disable-next-line vue/no-unused-vars -->
          <div v-for="(question, index) in survey.questions" :key="'q-' + question.id" class="question-container mb-4 p-3 border rounded">
            <h5>{{ question.text }} <span v-if="question.required" class="text-danger">*</span></h5>
            <p v-if="question.helpText" class="text-muted">{{ question.helpText }}</p>
            
            <!-- Single Choice Question -->
            <div v-if="question.type === 'single_choice'" class="form-group">
              <div v-for="option in question.options" :key="'opt-' + option.id" class="form-check">
                <input
                  type="radio"
                  class="form-check-input"
                  :name="'question-' + question.id"
                  :id="'option-' + question.id + '-' + option.id"
                  :value="option.id"
                  v-model="answers[question.id].optionId"
                  :required="question.required"
                >
                <label class="form-check-label" :for="'option-' + question.id + '-' + option.id">
                  {{ option.text }}
                  <img v-if="option.image" :src="option.image" class="option-image ml-2" :alt="option.text">
                </label>
              </div>
            </div>
            
            <!-- Multiple Choice Question -->
            <div v-else-if="question.type === 'multiple_choice'" class="form-group">
              <div v-for="option in question.options" :key="'opt-' + option.id" class="form-check">
                <input
                  type="checkbox"
                  class="form-check-input"
                  :id="'option-' + question.id + '-' + option.id"
                  :value="option.id"
                  v-model="answers[question.id].optionIds"
                  :required="question.required && answers[question.id].optionIds.length === 0"
                >
                <label class="form-check-label" :for="'option-' + question.id + '-' + option.id">
                  {{ option.text }}
                  <img v-if="option.image" :src="option.image" class="option-image ml-2" :alt="option.text">
                </label>
              </div>
            </div>
            
            <!-- Text Question -->
            <div v-else-if="question.type === 'text'" class="form-group">
              <textarea 
                class="form-control" 
                :id="'text-' + question.id" 
                :placeholder="question.placeholder || 'Enter your answer...'" 
                v-model="answers[question.id].textAnswer"
                :required="question.required"
                rows="3"
              ></textarea>
            </div>
            
            <!-- Rating Question -->
            <div v-else-if="question.type === 'rating'" class="form-group">
              <div class="rating-container d-flex justify-content-between">
                <div v-for="rating in ratingRange(question)" :key="'rate-' + rating" class="rating-item text-center">
                  <input
                    type="radio"
                    class="rating-input sr-only"
                    :name="'rating-' + question.id"
                    :id="'rating-' + question.id + '-' + rating"
                    :value="rating"
                    v-model="answers[question.id].ratingValue"
                    :required="question.required"
                  >
                  <label :for="'rating-' + question.id + '-' + rating" class="rating-label">
                    {{ rating }}
                  </label>
                </div>
              </div>
              <div class="d-flex justify-content-between mt-2">
                <small>{{ question.minValue }}</small>
                <small>{{ question.maxValue }}</small>
              </div>
            </div>
            
            <!-- Scale Question -->
            <div v-else-if="question.type === 'scale'" class="form-group">
              <input 
                type="range" 
                class="form-control-range" 
                :id="'scale-' + question.id" 
                :min="question.minValue" 
                :max="question.maxValue" 
                v-model.number="answers[question.id].scaleValue"
                :required="question.required"
              >
              <div class="d-flex justify-content-between">
                <small>{{ question.minValue }}</small>
                <small>{{ answers[question.id].scaleValue }}</small>
                <small>{{ question.maxValue }}</small>
              </div>
            </div>
            
            <!-- Date Question -->
            <div v-else-if="question.type === 'date'" class="form-group">
              <input 
                type="date" 
                class="form-control" 
                :id="'date-' + question.id" 
                v-model="dateInputs[question.id]"
                :required="question.required"
              >
            </div>
          </div>
          
          <div class="text-center mt-4">
            <button class="btn btn-primary btn-lg" @click.prevent="submitSurvey">Submit Answers</button>
          </div>
        </template>
        
        <template v-else>
          <h4 class="mb-4">Survey Results</h4>
          <div v-for="(result, index) in results" :key="index" class="mb-4">
            <h5>{{ result.text }}</h5>
            
            <!-- Choice question results -->
            <div v-if="result.optionResults && result.optionResults.length">
              <div class="list-group">
                <div v-for="optResult in result.optionResults" :key="optResult.optionId" 
                  class="list-group-item d-flex justify-content-between align-items-center">
                  {{ getOptionText(result.questionId, optResult.optionId) }}
                  <div>
                    <span class="badge badge-primary badge-pill mr-2">{{ optResult.count }}</span>
                    <span class="text-muted">{{ optResult.percentage.toFixed(1) }}%</span>
                  </div>
                </div>
              </div>
            </div>
            
            <!-- Rating or scale results -->
            <div v-else-if="result.averageRating || result.averageScale" class="mt-3">
              <div v-if="result.averageRating" class="alert alert-info">
                Average Rating: <strong>{{ result.averageRating.toFixed(1) }}</strong>
              </div>
              <div v-if="result.averageScale" class="alert alert-info">
                Average Scale Value: <strong>{{ result.averageScale.toFixed(1) }}</strong>
              </div>
              <h6 class="mt-3">Distribution:</h6>
              <div class="distribution-bars">
                <div v-for="(count, value) in (result.ratingCounts || result.scaleCounts)" :key="value"
                  class="d-flex align-items-center mb-2">
                  <span class="mr-2 text-right" style="width: 50px;">{{ value }}:</span>
                  <div class="progress flex-grow-1">
                    <div class="progress-bar" role="progressbar" 
                      :style="{ width: (count / result.totalVotes * 100) + '%' }"
                      :aria-valuenow="count" aria-valuemin="0" :aria-valuemax="result.totalVotes">
                      {{ count }}
                    </div>
                  </div>
                </div>
              </div>
            </div>
            
            <!-- Text answers -->
            <div v-else-if="result.textAnswers && result.textAnswers.length" class="mt-3">
              <h6>Responses ({{ result.textAnswers.length }}):</h6>
              <div class="list-group">
                <div v-for="(textResult, tIndex) in result.textAnswers" :key="'text-' + tIndex" 
                  class="list-group-item">
                  {{ textResult.answer }}
                  <span v-if="textResult.count > 1" class="badge badge-secondary ml-2">
                    {{ textResult.count }} times
                  </span>
                </div>
              </div>
            </div>
            
            <!-- Date answers -->
            <div v-else-if="result.dateDistribution" class="mt-3">
              <h6>Date Responses:</h6>
              <div class="list-group">
                <div v-for="(count, date) in result.dateDistribution" :key="date" 
                  class="list-group-item d-flex justify-content-between align-items-center">
                  {{ formatDate(date) }}
                  <span class="badge badge-primary badge-pill">{{ count }}</span>
                </div>
              </div>
            </div>
            
            <!-- Default case: show total votes -->
            <div v-else class="alert alert-secondary mt-3">
              Total responses: {{ result.totalVotes }}
            </div>
          </div>
        </template>
      </div>
      <div class="card-footer">
        <template v-if="showQuestions">
          <button type="button" class="btn btn-outline-secondary" @click="viewResults">View results</button>
        </template>
        <template v-else>
          <button type="button" class="btn btn-outline-primary" @click="showQuestions = true">Go back to survey</button>
        </template>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    survey: Object
  },
  data() {
    return {
      showQuestions: true,
      results: [],
      statusMessage: "",
      errorMessage: "",
      answers: {},
      dateInputs: {} // Store date inputs as YYYY-MM-DD strings
    };
  },
  created() {
    // Initialize answers object
    this.initializeAnswers();
  },
  methods: {
    initializeAnswers() {
      // Create empty answer objects for each question
      this.survey.questions.forEach(question => {
        this.answers[question.id] = {
          optionId: null,
          optionIds: [],
          textAnswer: "",
          ratingValue: null,
          scaleValue: question.minValue || 1,
          dateAnswer: null
        };
        
        // Initialize date inputs
        if (question.type === 'date') {
          this.dateInputs[question.id] = "";
        }
      });
    },
    ratingRange(question) {
      const min = question.minValue || 1;
      const max = question.maxValue || 5;
      return Array.from({length: max - min + 1}, (_, i) => min + i);
    },
    getOptionText(questionId, optionId) {
      const question = this.survey.questions.find(q => q.id === questionId);
      if (!question) return 'Unknown';
      
      const option = question.options.find(o => o.id === optionId);
      return option ? option.text : 'Unknown option';
    },
    formatDate(dateStr) {
      // Convert YYYY-MM-DD to more readable format
      try {
        const date = new Date(dateStr);
        return date.toLocaleDateString();
      } catch (e) {
        return dateStr;
      }
    },
    submitSurvey() {
      this.errorMessage = "";
      this.statusMessage = "";
      
      // Validate required fields
      let hasError = false;
      for (const question of this.survey.questions) {
        if (!question.required) continue;
        
        const answer = this.answers[question.id];
        
        if (question.type === 'single_choice' && answer.optionId === null) {
          hasError = true;
        } else if (question.type === 'multiple_choice' && answer.optionIds.length === 0) {
          hasError = true;
        } else if (question.type === 'text' && !answer.textAnswer.trim()) {
          hasError = true;
        } else if (question.type === 'rating' && answer.ratingValue === null) {
          hasError = true;
        } else if (question.type === 'date' && !this.dateInputs[question.id]) {
          hasError = true;
        }
      }
      
      if (hasError) {
        this.errorMessage = "Please fill in all required fields.";
        return;
      }
      
      // Process answers for submission
      const votes = [];
      
      for (const question of this.survey.questions) {
        const answer = this.answers[question.id];
        const baseVote = {
          survey: this.survey.id,
          question: question.id,
          timestamp: Math.floor(Date.now() / 1000)
        };
        
        switch (question.type) {
          case 'single_choice':
            if (answer.optionId !== null) {
              votes.push({
                ...baseVote,
                answerType: 'option',
                optionId: answer.optionId
              });
            }
            break;
            
          case 'multiple_choice':
            if (answer.optionIds.length > 0) {
              votes.push({
                ...baseVote,
                answerType: 'option',
                optionIds: answer.optionIds
              });
            }
            break;
            
          case 'text':
            if (answer.textAnswer.trim()) {
              votes.push({
                ...baseVote,
                answerType: 'text',
                textAnswer: answer.textAnswer.trim()
              });
            }
            break;
            
          case 'rating':
            if (answer.ratingValue !== null) {
              votes.push({
                ...baseVote,
                answerType: 'rating',
                ratingValue: answer.ratingValue
              });
            }
            break;
            
          case 'scale':
            if (answer.scaleValue !== null) {
              votes.push({
                ...baseVote,
                answerType: 'scale',
                scaleValue: answer.scaleValue
              });
            }
            break;
            
          case 'date':
            if (this.dateInputs[question.id]) {
              // Convert date string to timestamp
              const timestamp = new Date(this.dateInputs[question.id]).getTime() / 1000;
              votes.push({
                ...baseVote,
                answerType: 'date',
                dateAnswer: timestamp
              });
            }
            break;
        }
      }
      
      // Submit votes one by one
      const submitPromises = votes.map(vote => {
        return fetch('http://localhost:8082/vote', {
          method: "POST",
          body: JSON.stringify(vote),
          headers: { "Content-type": "application/json" }
        })
        .then(res => {
          if (!res.ok) throw new Error(`Error submitting vote: ${res.status}`);
          return res.json();
        });
      });
      
      Promise.all(submitPromises)
        .then(() => {
          this.statusMessage = "Thank you! Your responses have been recorded.";
          this.initializeAnswers(); // Reset form
          // Optionally show results after submission
          this.viewResults();
        })
        .catch(error => {
          this.errorMessage = "Failed to submit your responses. Please try again.";
          // eslint-disable-next-line no-console
          console.error(error);
        });
    },
    getResults() {
      this.errorMessage = "";
      this.statusMessage = "";

      // TODO: Make this URL controlled via env variable
      fetch(`http://localhost:8082/results/${this.survey.id}`)
        .then(res => {
          if (!res.ok) throw new Error(`Error fetching results: ${res.status}`);
          return res.json();
        })
        .then(data => {
          // Map question information to results
          this.results = this.survey.questions.map(question => {
            const questionResult = data.results.find(r => r.question === question.id) || {
              question: question.id,
              totalVotes: 0
            };
            
            return {
              ...questionResult,
              text: question.text,
              questionId: question.id,
              questionType: question.type
            };
          });
        })
        .catch(error => {
          this.errorMessage = "Cannot get survey results. Please try again.";
          // eslint-disable-next-line no-console
          console.error(error);
        });
    },
    viewResults() {
      this.getResults();
      this.showQuestions = false;
    }
  }
};
</script>

<style scoped>
.question-container {
  background-color: #ffffff;
  border-radius: 10px !important;
  box-shadow: 0 3px 10px rgba(0, 0, 0, 0.05);
  transition: transform 0.2s, box-shadow 0.2s;
  border: 1px solid #e9ecef !important;
}

.question-container:hover {
  transform: translateY(-2px);
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.1);
}

.option-image {
  max-height: 60px;
  max-width: 120px;
  object-fit: contain;
  border-radius: 5px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.rating-container {
  width: 100%;
  padding: 15px 0;
}

.rating-item {
  flex: 1;
  margin: 0 5px;
}

.rating-label {
  display: block;
  padding: 12px;
  cursor: pointer;
  border: 2px solid #dee2e6;
  border-radius: 8px;
  transition: all 0.3s;
  font-weight: 600;
  text-align: center;
}

.rating-input:checked + .rating-label {
  background-color: #4e73df;
  color: white;
  border-color: #4e73df;
  transform: scale(1.05);
}

.rating-label:hover {
  background-color: #e9ecef;
  transform: scale(1.05);
}

.progress {
  height: 25px;
  border-radius: 5px;
  margin-bottom: 5px;
}

.progress-bar {
  background-color: #4e73df;
  font-weight: 600;
}

.form-check {
  padding: 8px 10px;
  margin-bottom: 5px;
  border-radius: 5px;
  transition: background-color 0.2s;
}

.form-check:hover {
  background-color: #f8f9fa;
}

.form-check-input {
  margin-top: 0.2rem;
}

.form-check-label {
  font-weight: 500;
  cursor: pointer;
  display: flex;
  align-items: center;
}

.list-group-item {
  border-left: none;
  border-right: none;
  padding: 12px 20px;
  transition: background-color 0.2s;
}

.list-group-item:hover {
  background-color: #f8f9fa;
}

.card-footer {
  background-color: #f8f9fa;
  border-top: 1px solid #e9ecef;
  padding: 15px 20px;
}

.badge {
  padding: 6px 10px;
  font-weight: 600;
  border-radius: 20px;
}

.text-muted {
  color: #6c757d !important;
  font-size: 0.9rem;
}

.alert {
  border-radius: 8px;
  padding: 15px 20px;
}
</style>
