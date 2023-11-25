import { useState } from "react";
import dayjs, { Dayjs } from "dayjs";
import Box from "@mui/material/Box";
import Container from "@mui/material/Container";
import InputLabel from "@mui/material/InputLabel";
import MenuItem from "@mui/material/MenuItem";
import FormControl from "@mui/material/FormControl";
import Select, { SelectChangeEvent } from "@mui/material/Select";
import TopBar from "./TopBar";
import Grid from "@mui/material/Grid";
import Alert from "@mui/material/Alert";
import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";
import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs";
import { DatePicker } from "@mui/x-date-pickers/DatePicker";

function Grey() {
  const [alertOpen, setAlertOpen] = useState<Boolean>(false);
  const [alertMessage, setAlertMessage] = useState<string>("");
  const [symbol, setSymbol] = useState<string>("");
  const [date, setDate] = useState<Dayjs>(dayjs("2021-01-01"));
  const [expiry, setExpiry] = useState<string>("");

  const getExpiryList = (event: SelectChangeEvent) => {
    setSymbol(event.target.value);
    let symbol = event.target.value;
    fetch(
      `test.rebelscode.online/expiries/${symbol}/${date.format("YYYY-MM-DD")}`,
    ).then((response) => {
      if (response.status !== 200 || !response.ok) {
        console.log("error", date, symbol);
        setAlertOpen(true);
        setAlertMessage("Something Went Wrong!!!");
        return;
      }
      setAlertOpen(false);
      const data = response.json();
      console.log("success", data);
    });
  };

  const getOptionsChain = (event: SelectChangeEvent) => {
    setExpiry(event.target.value);
    console.log(symbol, expiry, event.target.value);
  };

  return (
    <div>
      <TopBar />
      {alertOpen && <Alert severity="error">{alertMessage}</Alert>}
      <Container maxWidth="xl" sx={{ mt: 2 }}>
        <Grid container alignItems="center" justify="center">
          <Grid item xs={12} md={4} sx={{ mt: 1 }}>
            <Box sx={{ display: "flex", justifyContent: "center" }}>
              <LocalizationProvider dateAdapter={AdapterDayjs}>
                <DatePicker
                  label="Start Date"
                  value={date}
                  onChange={(userDate) => setDate(userDate)}
                />
              </LocalizationProvider>
            </Box>
          </Grid>
          <Grid item xs={12} md={4}>
            <Box sx={{ display: "flex", justifyContent: "center" }}>
              <FormControl variant="standard" sx={{ m: 1, minWidth: 120 }}>
                <InputLabel id="symbol-label">Symbol</InputLabel>
                <Select
                  labelId="symbol-label"
                  value={symbol}
                  onChange={getExpiryList}
                  label="Symbol"
                >
                  <MenuItem value={"NIFTY"}>NIFTY</MenuItem>
                  <MenuItem value={"BANKNIFTY"}>BANKNIFTY</MenuItem>
                </Select>
              </FormControl>
            </Box>
          </Grid>
          <Grid item xs={12} md={4}>
            <Box sx={{ display: "flex", justifyContent: "center" }}>
              <FormControl variant="standard" sx={{ m: 1, minWidth: 120 }}>
                <InputLabel id="expiry-label">Expiry</InputLabel>
                <Select
                  labelId="expiry-label"
                  value={expiry}
                  onChange={getOptionsChain}
                  label="Expiry"
                >
                  <MenuItem value={"NIFTY"}>NIFTY</MenuItem>
                  <MenuItem value={"BANKNIFTY"}>BANKNIFTY</MenuItem>
                </Select>
              </FormControl>
            </Box>
          </Grid>
        </Grid>
      </Container>
    </div>
  );
}

export default Grey;
