// 点击关闭按钮隐藏公告
// $(document).ready(function () {
//   $(".cnetes").click(function () {
//     $(".title").hide();
//   });
// });

// vue实例
const con = {
  data() {
    return {
      message: "文章",
      text: '"欢迎来到⭐Rustling后台管理系统,此后台由JQ , Vue , JS等技术实现"',
      message1: "⭐Rustling后台管理系统",
      text1:
        "Rustling后台管理系统是一个基于Vue.js的后台管理系统, 它使用了最新的前端技术栈, 包括Vue.js, Element-UI, Vuex, Vue-Router, Axios等。",
    };
  },
};

Vue.createApp(con).mount("#con");
// 公告隐藏后向上自动滑动美化
$(document).ready(function () {
  const $title = $(".title");
  const $con = $("#con");
  // 原先是 30，现在改为 50（可按需调整，建议 40 - 80 之间）
  const extraUp = 20;
  $(".cnetes").click(function () {
    if ($title.is(":hidden")) {
      return;
    }

    const slideDistance = $title.outerHeight() || 50;
    const totalSlide = slideDistance + extraUp;
    // 总偏移量随之增大

    const conTop = $con.offset().top;
    const safeSlide = Math.min(totalSlide, conTop - 20); // 仍保留安全校验，防止滑过头

    $title.fadeOut(300);
    $con
      .css({
        position: "relative",
        top: 0,
      })
      .animate(
        {
          top: `-${safeSlide}px`,
        },
        500,
      );
  });
});

// 引入Vue3核心方法
const { createApp } = Vue;
