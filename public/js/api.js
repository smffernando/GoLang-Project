$(document).ready(function () {
$.ajax({
  type: "GET",
  url: '/catergory',
  success: function (data) {
    $.each(data.result, function (index, value) {
      $("#cdropdown").append("<option value=" + value.Id + ">" + value.Name + "</option>");
    });
  }
});
$("#cdropdown").on('change', function () {
  console.log("list item selected");
  var val = $(this).val();
  console.log(val);
  $.ajax({
      type: "GET",
      url: '/productsubcategory/'+val,
      success: function (data) {
          $("#sdropdown").html("");
          $.each(data.result, function (index, value) {
              $("#sdropdown").append("<option value=" + value.Id + ">" + value.Name + "</option>");
          });
      }
  });
});

$.ajax({
  type: "GET",
  url: '/brand',
  success: function (data) {
    $.each(data.result, function (index, value) {
      $("#drpdown").append("<option value=" + value.Id + ">" + value.Name + "</option>");
    });
  }
});

// get products
$.ajax({
  type: "GET",
  url: '/allproduct/1',
  success: function (data) {
    $.each(data.result, function (index, value) {
      $("#results_list").append("<a href='good' class='service_search_panel_link link'><div class='row summary_row wow fadeIn'><div class='col-4 no_space_right summary_image_column'><img src='https://111kuroyanagi1.files.wordpress.com/2013/02/1996-nissan-atlas-200-tipper-truck-sale-japan-60k.jpg' class='summary_image img-fluid'/></div><div class='col-8'><p class='summary_text' id='summary_title'>"+value.Name+"</p><p class='summary_text' id='summary_city'><i class='fas fa-map-marker-alt icon_type2'></i>"+value.Headder+"</p><p class='summary_text' id='summary_price'> <i class='fas fa-tag icon_type2'></i>"+value.Description+"</p><button class='summary_book_now_button'> Book Now </button></div></div></a>");
    });
    // set product images
    // set_images(data.productImages)
    // load pagination
    set_pagination(data.pages)
    set_images(data.productImages)
  }
});

// pagination
$("#pagination").on('click','.pagination_page',function(){
  var page=$(this).attr('data-pagenumber');
  $.ajax({
    type: "GET",
    url: '/allproduct/'+page,
    success: function (data) {
      $("#results_list").html("");
      $.each(data.result, function (index, value) {
        $("#results_list").append("<a href='good/' class='service_search_panel_link link'><div class='row summary_row wow fadeIn'><div class='col-4 no_space_right summary_image_column'><img src='https://111kuroyanagi1.files.wordpress.com/2013/02/1996-nissan-atlas-200-tipper-truck-sale-japan-60k.jpg' class='summary_image img-fluid'/></div><div class='col-8'><p class='summary_text' id='summary_title'>"+value.Name+"</p><p class='summary_text' id='summary_city'><i class='fas fa-map-marker-alt icon_type2'></i>"+value.Headder+"</p><p class='summary_text' id='summary_price'> <i class='fas fa-tag icon_type2'></i>"+value.Description+"</p><button class='summary_book_now_button'> Book Now </button></div></div></a>");
      });
    // set product images
    set_images(data.productImages)
    }
  }); 
});

// set product images
function set_images(imagelist){
    $.each(imagelist, function (index, value) {
      console.log(value);
    });
}

// set pagination
function set_pagination(page_count){
      $("#pagination").append("<a href='#'>First<<</a>");
   for(i=1; i<=page_count+1; i++){
      $("#pagination").append("<a href='#' class='pagination_page' data-pagenumber="+i+">"+i+"</a>")
   }
    $("#pagination").append("<a href=#''>>>Last</a>")
}

});
