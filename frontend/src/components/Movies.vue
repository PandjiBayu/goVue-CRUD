<template>
  <div class="container">
    <div class="d-flex justify-content-between mb-3 ml-3">
      <h3>DATA MOVIES</h3>

      <div class="mr-3">
        <button class="btn btn-success" 
        v-on:click="addMovie()">Add Movie</button>
      </div>
    </div>

    <div v-if="message" class="alert alert-success">{{ this.message }}</div>
    <div class="container">
      <table class="table table-striped">
        <thead>
          <tr> 
            <th>Movie's Title</th>
            <th>Genre</th>
            <th>Release Year</th>
            <th>Update</th>
            <th>Delete</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="movie in movies" :key="movie._id">
            <td>{{ movie.title }}</td>
            <td>{{ movie.genre }}</td>
            <td>{{ movie.year }}</td>
            <td>
              <button class="btn btn-warning" 
              v-on:click="updateMovie(movie._id)">
                Update
              </button>
            </td>
            <td>
              <button class="btn btn-danger" 
              v-on:click="deleteMovie(movie._id)">
                Delete
              </button>
            </td>
          </tr>
        </tbody>
      </table>

    </div>
  </div>
</template>

<script>
import movieDataService from "../service/movieDataService";

export default {
  name: "Movies",
  data() {
    return {
      movies: [],
      message: "",
    };
  },
  methods: {
    refreshMovies() {
        movieDataService.getAllMovies().then((res) => {
          this.movies = res.data.data.data;
          console.log(res.data.data.data);
        }).catch(error => {
            console.log(error);
        });
    },
    addMovie() {
        this.$router.push({name: 'Create'});
    },
    updateMovie(id) {
      this.$router.push(`/movie/${id}`);
    },
    deleteMovie(id) {
      if(!confirm("Are you sure?")){
            return;
        }
      movieDataService.deleteMovie(id)
      .then(() => {
        this.refreshMovies();
      });
    },
    
  },
  created() {
    this.refreshMovies();
  },
  
};
</script>