$(document).ready(function(){

// $('#register').on('click',function(){
    $("#userform").on('submit',function(){
    data = $(this).serialize();

    $.ajax({
        type:'POST',
        url: 'http://74.208.71.132:3000/newuser',
        data: data,
        dataType:'json',
        success: function(newRegister){
            $orders.append('<li>first_name: '+ newRegister.first_name+',last_name:'+ newRegister.last_name+',email:'+ newRegiste.email +',password:'+newRegiste.password+'</li>');
        }
        // error: function(){
        //     alert('Error Supplier Register');

        // }
    });
    // window.location.replace("supplier.html");
return false;

});

});

