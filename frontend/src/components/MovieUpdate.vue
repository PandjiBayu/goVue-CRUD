<template>
  <div>
    <h3>Movie</h3>
    <div class="container">
      <form @submit="validateAndSubmit">
        <div v-if="errors.length">
          <div
            class="alert alert-danger"
            v-bind:key="index"
            v-for="(error, index) in errors"
          >
            {{ error }}
          </div>
        </div>
        <fieldset class="form-group">
          <label>Movie's Title</label>
          <input type="text" class="form-control"
           v-model="title" />
        </fieldset>
        <fieldset class="form-group">
          <label>Genre</label>
          <input type="text" class="form-control"
           v-model="genre" />
        </fieldset>
        <fieldset class="form-group">
          <label>Release Year</label>
          <input type="text" class="form-control" 
          v-model="year" />
        </fieldset>
        <button class="btn btn-success" 
        type="submit">Update</button>
      </form>
    </div>
  </div>
</template>
<script>

import movieDataService from "../service/movieDataService";

export default {
  name: "MovieUpdate",
  data() {
    return {
      title: "",
      genre: "",
      year: "",
      errors: [],
    };
  },
  computed: {
    id() {
      return this.$route.params.id;
    },
  },
  methods: {
    refreshMovieDetails() {
      movieDataService.getIdMovie(this.id).then((res) => {
        this.title = res.data.data.data.title;
        this.genre = res.data.data.data.genre;
        this.year = res.data.data.data.year;
      });
    },
    validateAndSubmit(e) {
      e.preventDefault();
      this.errors = [];
      if (!this.title) {
        this.errors.push("Enter valid values");
      } else if (this.title.length < 2) {
        this.errors.push
        ("Enter atleast 3 characters in Title");
      }
      if (!this.genre) {
        this.errors.push("Enter valid values");
      } else if (this.genre.length < 4) {
        this.errors.push
        ("Enter atleast 4 characters in Genre");
      }
      if (!this.year) {
        this.errors.push("Enter valid values");
      } else if (this.year.length < 4) {
        this.errors.push
        ("Enter correct year");
      }
      if (this.errors.length === 0) {
        console.log('123');
          movieDataService.updateMovie(this.id, {
            id: this.id,
            title: this.title,
            genre: this.genre,
            year: this.year,
          }).then(() => {
            this.$router.push({name: 'Movies'});
          });
      }
    },
  },
  mounted:function() {
    this.refreshMovieDetails();
  },
};
</script>
