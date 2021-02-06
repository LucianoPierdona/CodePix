import {
  Body,
  Controller,
  Get,
  HttpCode,
  Inject,
  InternalServerErrorException,
  NotFoundException,
  Param,
  ParseUUIDPipe,
  Post,
  Query,
  UnprocessableEntityException,
  ValidationPipe,
} from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { BankAccount } from '../../models/bank-account.model';
import { PixKey } from '../../models/pix-key.model';
import { Repository } from 'typeorm';

@Controller('bank-account/:bankAccountId/pix-keys')
export class PixKeyController {
  constructor(
    @InjectRepository(PixKey)
    private pixKeyRepo: Repository<PixKey>,
  ) {}

  @Get()
  index(
    @Param('bankAccountId', new ParseUUIDPipe({ version: '4' }))
    bankAccountId: string,
  ) {
    return this.pixKeyRepo.find({
      where: {
        bank_account_id: bankAccountId,
      },
      order: {
        created_at: 'DESC',
      },
    });
  }
}
