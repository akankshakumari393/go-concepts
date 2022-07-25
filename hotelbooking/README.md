


hotelId hotelname




10 MMT   ---   10 serviceA
               






serviceName serviceID





serviceId HotelID
1   2
1   3
2   1


                                                                         {
                                                                             CN
                                                                         }
PaymentInfo{
    cardNumber string
    CVV 
    bookingAmount
    firstName
    lastName
    pendingAmount
    inventory
}


type Service interface {
    CreateRequest
    SendRequest
    ReadResponse
}

func (*service1) SendRequest() Response {

}

func (*service1) CreateRequest(PaymentInfo) Response {

}




func postBooking(hotelId, paymentInfo) {
    // get the list of service exposing the hotel 
         services := getAllServices(hotelId)

    for _, service := range service {

        go func() {
        Request := service.CreateRequest(paymentInfo)
        service.SendRequest()




        }




    }



    serviceA --
    {CN = carDNumber}

}







table 

        cardNumber CVV bookingAmount firstName LastName
serviceA  CN
ServiceB




