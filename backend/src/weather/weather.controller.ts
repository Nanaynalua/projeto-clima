import { Controller, Get, Post, Body, Param } from '@nestjs/common';
import { WeatherService } from './weather.service';
import { Weather } from './weather.schema';

@Controller('weather')
export class WeatherController {
  constructor(private readonly weatherService: WeatherService) {}

  // POST /weather → cria um registro
  @Post()
  async create(@Body() data: { city: string; temperature: number; humidity: number }): Promise<Weather> {
    return this.weatherService.create(data);
  }

  // GET /weather → lista todos os registros
  @Get()
  async findAll(): Promise<Weather[]> {
    return this.weatherService.findAll();
  }

  // GET /weather/:city → lista registros de uma cidade
  @Get(':city')
  async findByCity(@Param('city') city: string): Promise<Weather[]> {
    return this.weatherService.findByCity(city);
  }
}
