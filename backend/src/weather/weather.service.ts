import { Injectable } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { Model } from 'mongoose';
import { Weather } from './weather.schema';

@Injectable()
export class WeatherService {
  constructor(
    @InjectModel(Weather.name) private weatherModel: Model<Weather>,
  ) {}

  // Criar um novo registro de clima
  async create(data: {
    city: string;
    temperature: number;
    humidity: number;
    timestamp?: Date;
  }): Promise<Weather> {
    const createdWeather = new this.weatherModel({
      ...data,
      timestamp: data.timestamp || new Date(),
    });
    return createdWeather.save();
  }

  // Buscar todos os registros
  async findAll(): Promise<Weather[]> {
    return this.weatherModel.find().exec();
  }

  // Buscar registros por cidade
  async findByCity(city: string): Promise<Weather[]> {
    return this.weatherModel.find({ city }).exec();
  }

  // Deletar todos os registros (útil para testes)
  async clear(): Promise<void> {
    await this.weatherModel.deleteMany({});
  }
}
