<!DOCTYPE html>
<html>
<head>
	<title>Vue Axios Example</title>
	<script src="https://cdnjs.cloudflare.com/ajax/libs/vue/2.4.2/vue.js"></script>
	<script src="https://unpkg.com/axios/dist/axios.min.js"></script>
</head>
<body>
<div id="app">
  <ul>
    <li v-for="movie in movies">{{ movie.title }}</li>
  </ul>
</div>

<script>
	
	new Vue({
  el: '#app',
  data: {
    movies: []
  },
  created () {
    var vm = this
    axios.get('https://jsonplaceholder.typicode.com/users')
      .then(function (response) {
        console.log(response)
        vm.movies = response.data.data
      })
  }
})

</script>

</body>
</html>