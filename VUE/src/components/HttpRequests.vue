<template>
    <div>
        <input type="text" v-model="user.nombre">
        <button @click="sendStr">Send String</button>
        <hr>
        <button @click="getStr">Get String</button>
        <hr>
        <input type="text" v-model="user.nombre">
        <br>
        <input type="text" v-model="user.edad">
        <br>
        <input type="text" v-model="user.coche">
        <br>
        <input type="text" v-model="user.ciudad">
        <br>
        <button @click="sendJson">Send Json</button>
        <hr>
        <h2>{{user.nombre}}</h2>
        <br>
        <h2>{{user.edad}}</h2>
        <br>
        <h2>{{user.coche}}</h2>
        <br>
        <h2>{{user.ciudad}}</h2>
        <br>
        <button @click="getJson">Get Json</button>
        <br>
        <button @click="sendJsonArray">Send Json Array</button>
        <br>
        <button @click="getJsonArray">Get Json Array</button>
        <ul>
            <li v-for="u in users">{{u.nombre}} - {{u.edad}} - {{u.coche}} - {{u.ciudad}}</li>
        </ul>
        <hr>
        <button @click="getStrArray">Get String Array</button>
        <hr>
        <button @click="sendStrArray">Send String Array</button>
        <br>
    </div>
</template>

<script>
export default {
  data() {
      return {
        user: {
            nombre: '',
            edad: '',
            coche: '',
            ciudad: ''
        },
        users: []
      }
  },
  methods: {
      sendStr(){
          this.$http.post('http://localhost:5000/vueb', this.user.nombre)
      },
      getStr(){
        this.$http.get('http://localhost:5000/vuea').then(response => {
            console.log(response.body)
        });
      },
      sendStrArray(){
            const xx = ["Isabel", "Nuria", "Laura", "Natalia"];
            const yy = xx.toString();
            this.$http.post('http://localhost:5000/vueh', yy);
      },
      getStrArray(){
        this.$http.get('http://localhost:5000/vueg').then(response => {
            let xx = response.body;
            let yy = xx.split(",");
            for(let x in yy){
                console.log(yy[x]);
            }
        });
      },
      sendJsonArray(){
        const xx = [{"nombre": "Isabel", "edad": "32", "coche": "BMW", "ciudad": "Santander"},{"nombre": "Nuria", "edad": "40", "coche": "Ferrari", "ciudad": "Oviedo"}];
        this.$http.post('http://localhost:5000/vuef', xx);
      },
      getJsonArray(){
          this.$http.post('http://localhost:5000/vuec').then(response => {
              return response.json()

          }).then(data => {
              console.log(data);
              const resultArray = [];
              for(let key in data){
                  resultArray.push(data[key]);
              }
              this.users = resultArray;
          });
      },
      sendJson(){
        this.$http.post('http://localhost:5000/vuee', this.user)
      },
      getJson(){
        this.$http.post('http://localhost:5000/vued').then(response => {
            return response.json()
        }).then(data => {
            this.user.nombre = data["nombre"];
            this.user.edad = data["edad"];
            this.user.coche = data["coche"];
            this.user.ciudad = data["ciudad"];
        });
      }
  }
}
</script>

<style lang="scss">

</style>
