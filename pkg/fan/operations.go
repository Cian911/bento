package fan

const (
  HEADER = "FDFD"
  DEVICE_TYPE = "02"

  // Unit Operations
  OP_READ_REQUEST = "01"
  OP_WRITE_REQUEST = "02"
  OP_WRITE_RETURN_REQUEST = "03"
  OP_INC_REQUEST = "04"
  OP_DEC_REQUEST = "05"

  // Turn on/off fan
  OP_ON = "01"
  OP_OFF = "00"

  // Airflow Operations
  OP_AIR_OUT = "00"
  OP_AIR_INVERT = "01"
  OP_AIR_IN = "02"

  // Airflow Operation Request
  OP_AIRFLOW_REQUEST = "00B7"

  // Unit Operation Request
  OP_UNIT_OPERATION_REQUEST = "0001"

  // Temperature
  OP_AIR_TEMP_REQUEST = "0020"

  // Wifi
  OP_WIFI_CLIENT_NAME_REQUEST = "0095"
  OP_WIFI_OPERATION_REQUEST = "0094"
  OP_WIFI_STATUS_REQUEST = "00A1"
  OP_WIFI_IP_ASSIGNED_REQUEST = "00A3"

  // Speed
  OP_SPEED_MODE_REQUEST = "0002"
)
