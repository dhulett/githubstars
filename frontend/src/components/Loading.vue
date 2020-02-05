<template>
  <div>
    <div id="outer-barG">
      <div id="front-barG" class="bar-animationG">
        <div id="barG_1" class="bar-lineG"></div>
        <div id="barG_2" class="bar-lineG"></div>
        <div id="barG_3" class="bar-lineG"></div>
      </div>
    </div>
    <div>{{ loadingText }}</div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      loadingText: "Getting the repositories list from Github..."
    };
  }
};
</script>

<style scoped>
#outer-barG {
  height: 20px;
  width: 100px;
  border: 2px solid rgba(0, 0, 0, 0.84);
  overflow: hidden;
  background-color: rgba(255, 255, 255, 0.5);
  margin-top: 200px;
  margin-left: 600px;
}

.bar-lineG {
  background-color: rgba(104, 98, 98, 0.2);
  float: left;
  width: 14px;
  height: 93px;
  margin-right: 19px;
  margin-top: -22px;
  transform: rotate(45deg);
  -o-transform: rotate(45deg);
  -ms-transform: rotate(45deg);
  -webkit-transform: rotate(45deg);
  -moz-transform: rotate(45deg);
}

.bar-animationG {
  margin-left: 130px;
  width: 130px;
  animation-name: bar-animationG;
  -o-animation-name: bar-animationG;
  -ms-animation-name: bar-animationG;
  -webkit-animation-name: bar-animationG;
  -moz-animation-name: bar-animationG;
  animation-duration: 4.58s;
  -o-animation-duration: 4.58s;
  -ms-animation-duration: 4.58s;
  -webkit-animation-duration: 4.58s;
  -moz-animation-duration: 4.58s;
  animation-iteration-count: infinite;
  -o-animation-iteration-count: infinite;
  -ms-animation-iteration-count: infinite;
  -webkit-animation-iteration-count: infinite;
  -moz-animation-iteration-count: infinite;
  animation-direction: normal;
  -o-animation-direction: normal;
  -ms-animation-direction: normal;
  -webkit-animation-direction: normal;
  -moz-animation-direction: normal;
}
/*
@keyframes bar-animationG {
  0% {
    margin-left: 120px;
    margin-top: -14px;
  }

  100% {
    margin-left: -99px;
    margin-top: -14px;
  }
}

@-o-keyframes bar-animationG {
  0% {
    margin-left: 120px;
    margin-top: -14px;
  }

  100% {
    margin-left: -99px;
    margin-top: -14px;
  }
}

@-ms-keyframes bar-animationG {
  0% {
    margin-left: 120px;
    margin-top: -14px;
  }

  100% {
    margin-left: -99px;
    margin-top: -14px;
  }
}

@-webkit-keyframes bar-animationG {
  0% {
    margin-left: 120px;
    margin-top: -14px;
  }

  100% {
    margin-left: -99px;
    margin-top: -14px;
  }
}

@-moz-keyframes bar-animationG {
  0% {
    margin-left: 120px;
    margin-top: -14px;
  }

  100% {
    margin-left: -99px;
    margin-top: -14px;
  }
} */

h2 {
  text-align: center;
}
</style>
