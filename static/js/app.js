//----JS Module for Modals & Search availability modal----
function Prompt() {
  //toast in top right corner
  function successToast(c) {
    const { msg = "", icon = "success" } = c;

    const Toast = Swal.mixin({
      toast: true,
      position: "top-end",
      showConfirmButton: false,
      timer: 3000,
      timerProgressBar: true,
      didOpen: (toast) => {
        toast.addEventListener("mouseenter", Swal.stopTimer);
        toast.addEventListener("mouseleave", Swal.resumeTimer);
      },
    });

    Toast.fire({
      icon: icon,
      title: msg,
    });
  }
  //error modal
  function error(c) {
    const { title = "", msg = "", icon = "error", buttonText = "Ok" } = c;

    Swal.fire({
      title: title,
      text: msg,
      icon: icon,
      confirmButtonText: buttonText,
    });
  }
  //modal which can display html form & process data
  async function custom(c) {
    const { icon = "", msg = "", title = "", showConfirmButton = true } = c;
    //display modal and return result data from possible html form
    const { value: result } = await Swal.fire({
      icon: icon,
      title: title,
      html: msg,
      focusConfirm: true,
      backdrop: false,
      showCancelButton: true,
      showConfirmButton: showConfirmButton,
      willOpen: () => {
        //willOpen function can be passed to custom function
        if (c.willOpen !== undefined) {
          c.willOpen();
        }
      },
      didOpen: () => {
        if (c.didOpen !== undefined) {
          c.didOpen();
        }
      },
    });
    //->call custom callback function given by calling function in params with result data
    if (result) {
      //Swal.fire(JSON.stringify(result));
      if (result.dismiss !== Swal.DismissReason.cancel) {
        if (result.value !== "") {
          if (c.callback !== undefined) {
            c.callback(result);
          } else {
            c.callback(false);
          }
        } else {
          c.callback(false);
        }
      }
    }
  }

  return {
    successToast: successToast,
    error: error,
    custom: custom,
  };
}
