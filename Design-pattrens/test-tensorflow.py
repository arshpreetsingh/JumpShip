import yfinance as yf
import pandas as pd
import numpy as np
from sklearn.preprocessing import MinMaxScaler
from tensorflow.keras.models import Sequential
from tensorflow.keras.layers import LSTM, Dense
import matplotlib.pyplot as plt

# define function to get stock data
def get_stock_data(ticker):
    stock_data = yf.Ticker(ticker)
    stock_df = stock_data.history(period="max")
    return stock_df

# get stock data for 'AAPL' and 'MSFT'
df_aapl = get_stock_data("AAPL")
df_msft = get_stock_data("MSFT")

# define function to prepare data for LSTM model
def prepare_data_for_lstm(data):
    # create numpy array
    dataset = np.array(data)
    # scale the data
    scaler = MinMaxScaler(feature_range=(0,1))
    scaled_data = scaler.fit_transform(dataset.reshape(-1,1))
    # create training data and labels
    x_train = []
    y_train = []
    for i in range(60,len(scaled_data)):
        x_train.append(scaled_data[i-60:i,0])
        y_train.append(scaled_data[i,0])
    # convert the lists to numpy arrays
    x_train = np.array(x_train)
    y_train = np.array(y_train)
    # reshape the data for LSTM model
    x_train = np.reshape(x_train, (x_train.shape[0], x_train.shape[1], 1))
    return x_train, y_train, scaler

# prepare data for LSTM model
x_train_aapl, y_train_aapl, scaler_aapl = prepare_data_for_lstm(df_aapl["Close"])
x_train_msft, y_train_msft, scaler_msft = prepare_data_for_lstm(df_msft["Close"])

# define function to create LSTM model
def create_lstm_model(x_train):
    model = Sequential()
    model.add(LSTM(units=50, return_sequences=True, input_shape=(x_train.shape[1],1)))
    model.add(LSTM(units=50, return_sequences=False))
    model.add(Dense(units=25))
    model.add(Dense(units=1))
    return model

# create LSTM model for AAPL
model_aapl = create_lstm_model(x_train_aapl)

# compile the model
model_aapl.compile(optimizer='adam', loss='mean_squared_error')

# train the model
model_aapl.fit(x_train_aapl, y_train_aapl, epochs=1, batch_size=1)

# create LSTM model for MSFT
model_msft = create_lstm_model(x_train_msft)

# compile the model
model_msft.compile(optimizer='adam', loss='mean_squared_error')

# train the model
model_msft.fit(x_train_msft, y_train_msft, epochs=1, batch_size=1)

# define function to make predictions
def make_predictions(model, x_train, scaler):
    # make predictions
    predictions = model.predict(x_train)
    # undo scaling
    predictions = scaler.inverse_transform(predictions)
    return predictions

# make predictions for AAPL
predictions_aapl = make_predictions(model_aapl, x_train_aapl, scaler_aapl)

# make predictions for MSFT
predictions_msft = make_predictions(model_msft, x_train_msft, scaler_msft)

# plot actual vs predicted prices for AAPL
plt.plot(df_aapl["Close"].values, label="Actual Price")
plt.plot(predictions_aapl, label="Predicted Price")
plt.title("AAPL Actual vs Predicted Prices")
plt.xlabel("Time")
plt.ylabel("Price")
plt.legend()
plt.show()

# plot actual vs predicted prices
