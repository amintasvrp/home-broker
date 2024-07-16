import { Body, Controller, Post } from '@nestjs/common';
import { AssetsService } from './assets.service';
import { Asset } from '@prisma/client';

@Controller('assets')
export class AssetsController {
  constructor(private readonly assetsService: AssetsService) {}

  @Post()
  create(
    @Body() body: { id: string; symbol: string; price: number },
  ): Promise<Asset> {
    return this.assetsService.create(body);
  }
}
