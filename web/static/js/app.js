import $ from "jquery"
import Turbolinks from "turbolinks";

require("jquery");
require("bootstrap-datepicker");
require("js-cookie");

$(document).ready(function(){
    Turbolinks.start();

    // Update profile image.
    let imageUrl = $('.profile-image').attr('rel');
    $('.profile-image-thumb, .profile-image').attr('src', imageUrl);

    // Datepicker
    if($('#dashboardDate').length) {
        let date = new Date();
        let today = new Date(date.getFullYear(), date.getMonth(), date.getDate());
        $('#dashboardDate').datepicker({
            format: "dd-MM-yyyy",
            todayHighlight: true,
            autoclose: true
        });
        $('#dashboardDate').datepicker('setDate', today);
    }

    // Logout
    $(document).ready(function() {
        $('.log-out').click(function(e) {
            Cookies.remove('auth-session');
        });
    });
});
