const vm = Vue.createApp({
  data() {
    return {
      message: "<b>Hello, Vue 3!</b>",
      a: 2,
      b: 3,
      name: "",
      url: "https://vuejs.org/",
    };
  },
  methods: {
    clickEventHandle() {
      alert("Cliked");
    },
    onSubmit() {
      alert("hello," + this.name);
    },
  },
});

vm.mount("#app");
