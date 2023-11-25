import { createTheme } from "@mui/material/styles";
import { grey } from "@mui/material/colors";

const theme = createTheme({
  palette: {
    primary: {
      main: grey[500],
    },
    secondary: {
      main: grey[200],
    },
    // background: {
    //   default: "silver",
    // },
  },
});

export default theme;
