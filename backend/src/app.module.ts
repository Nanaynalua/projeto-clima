import { Module } from '@nestjs/common';
import { MongooseModule } from '@nestjs/mongoose';
import { WeatherModule } from './weather/weather.module';

@Module({
  imports: [
    MongooseModule.forRoot('mongodb://mongodb:27017/projeto-clima'),
    WeatherModule,
  ],
})
export class AppModule {}
