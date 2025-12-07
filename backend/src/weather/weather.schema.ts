import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { Document } from 'mongoose';

@Schema()
export class Weather extends Document {
  @Prop()
  city: string;

  @Prop()
  temperature: number;

  @Prop()
  humidity: number;

  @Prop()
  timestamp: Date;
}

export const WeatherSchema = SchemaFactory.createForClass(Weather);
