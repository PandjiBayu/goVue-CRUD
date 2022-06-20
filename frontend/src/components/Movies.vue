<template>
  <div class="container">
    <h3>DATA MOVIES</h3>

    <div v-if="message" class="alert alert-success">{{ this.message }}</div>
    <div class="container">
      <table class="table table-striped">
        <thead>
          <tr> 
            <th>Movie's Title</th>
            <th>Genre</th>
            <th>Year Release</th>
            <th>Update</th>
            <th>Delete</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="movie in movies" :key="movie.id">
            <td>{{ movie.title }}</td>
            <td>{{ movie.genre }}</td>
            <td>{{ movie.year }}</td>
            <td>
              <button class="btn btn-warning" 
              v-on:click="updateMovie(movie.id)">
                Update
              </button>
            </td>
            <td>
              <button class="btn btn-danger" 
              v-on:click="deleteMovie(movie.id)">
                Delete
              </button>
            </td>
          </tr>
        </tbody>
      </table>

      <div class="row">
        <button class="btn btn-success" 
        v-on:click="addMovie()">Add Movie</button>
      </div>
    </div>

    <!-- modal -->
    <!-- <div class="modal fade" id="exampleModal" tabindex="-1" aria-labelledby="exampleModalLabel" aria-hidden="true">
      <div class="modal-dialog modal-lg modal-dialog-centered">
        <div class="modal-content">
          <div class="modal-header">
              <h5 class="modal-title" id="exampleModalLabel">{{modalTitle}}</h5>
              <button type="button" class="btn-close" data-bs-dismiss="modal"
              aria-label="Close"></button>
          </div>

          <div class="modal-body">
              <div class="p-2 w-50 bd-highlight">
                  <div class="input-group mb-3">
                      <span class="input-group-text">Movie's Title</span>
                      <input type="text" class="form-control" v-model="title">
                  </div>

                  <div class="input-group mb-3">
                      <span class="input-group-text">Genre</span>
                      <input type="text" class="form-control" v-model="genre">
                  </div>

                  <div class="input-group mb-3">
                      <span class="input-group-text">Year Release</span>
                      <input type="date" class="form-control" v-model="year">
                  </div>
              </div>

              <button type="button" @click="createClick()"
              v-if="id==0" class="btn btn-primary">
              Create
              </button>

              <button type="button" @click="updateClick()"
              v-if="id!=0" class="btn btn-primary">
              Update
              </button>
          </div>
        </div>
      </div>
    </div> -->
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
       movieDataService.getAllMovies(this.movies).then((res) => {
        this.movies = res.data;
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
    // createClick(){
    //     movieDataService.createMovie({
    //         title:this.title,
    //         genre:this.genre,
    //         year:this.year
    //     })
    //     .then(()=>{
    //         this.refreshMovies();
    //         this.$router.push("/movies")
    //     });
    // },
    // updateClick(){
    //     movieDataService.updateMovie(this.id,{
    //         id:this.id,
    //         title:this.title,
    //         genre:this.genre,
    //         year:this.year
    //     })
    //     .then((response)=>{
    //         this.refreshMovies();
    //         this.$router.push("/movies")
    //     });
    // },
  },
  created() {
    this.refreshMovies();
  },
};
</script>