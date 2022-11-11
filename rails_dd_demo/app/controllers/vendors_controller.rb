class VendorsController < ApplicationController
  def index
    data = {
      vendors: [{
        id: 1,
        name: "MCD"
      }]
    }

    render json: data
  end
end
